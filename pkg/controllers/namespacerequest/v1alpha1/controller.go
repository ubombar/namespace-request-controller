package v1alpha1

import (
	"context"
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"

	"github.com/ubombar/namespace-request-controller/pkg/apis/namespacerequest/v1alpha1"
	clientset "github.com/ubombar/namespace-request-controller/pkg/generated/clientset/versioned"
	samplescheme "github.com/ubombar/namespace-request-controller/pkg/generated/clientset/versioned/scheme"
	informers "github.com/ubombar/namespace-request-controller/pkg/generated/informers/externalversions/namespacerequest/v1alpha1"
	listers "github.com/ubombar/namespace-request-controller/pkg/generated/listers/namespacerequest/v1alpha1"
)

const controllerAgentName = "namespacerequest-controller"

const (
	// SuccessSynced is used as part of the Event 'reason' when a Foo is synced
	successSynced = "Synced"
	// ErrResourceExists is used as part of the Event 'reason' when a Foo fails
	// to sync due to a Deployment of the same name already existing.
	errResourceExists = "ErrResourceExists"

	// MessageResourceExists is the message used for Events when a resource
	// fails to sync due to a Deployment already existing
	messageResourceExists = "Resource %q already exists and is not managed by Namespace Request"
	// MessageResourceSynced is the message used for an Event fired when a Foo
	// is synced successfully
	messageResourceSynced = "Namespace Request synced successfully"
)

// Controller is the controller implementation for Foo resources
type Controller struct {
	// kubeclientset is a standard kubernetes clientset
	kubeclientset kubernetes.Interface
	// ubombarclientset is a clientset for our own API group
	ubombarclientset clientset.Interface

	namespacesLister  corelisters.NamespaceLister
	namespacessSynced cache.InformerSynced

	namespaceRequestLister listers.NamespaceRequestLister
	namespaceRequestSynced cache.InformerSynced

	// workqueue is a rate limited work queue. This is used to queue work to be
	// processed instead of performing it as soon as a change happens. This
	// means we can ensure we only process a fixed amount of resources at a
	// time, and makes it easy to ensure we are never processing the same item
	// simultaneously in two different workers.
	workqueue workqueue.RateLimitingInterface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	recorder record.EventRecorder
}

// NewController returns a new sample controller
func NewController(
	ctx context.Context,
	kubeclientset kubernetes.Interface,
	ubombarclientset clientset.Interface,
	namespaceInformer coreinformers.NamespaceInformer,
	namespaceRequestInformer informers.NamespaceRequestInformer) *Controller {
	logger := klog.FromContext(ctx)

	// Create event broadcaster
	// Add namespace-request-controller types to the default Kubernetes Scheme so Events can be
	// logged for namespace-request-controller types.
	utilruntime.Must(samplescheme.AddToScheme(scheme.Scheme))
	logger.V(4).Info("Creating event broadcaster")

	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartStructuredLogging(0)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

	controller := &Controller{
		kubeclientset:          kubeclientset,
		ubombarclientset:       ubombarclientset,
		namespacesLister:       namespaceInformer.Lister(),
		namespacessSynced:      namespaceInformer.Informer().HasSynced,
		namespaceRequestLister: namespaceRequestInformer.Lister(),
		namespaceRequestSynced: namespaceRequestInformer.Informer().HasSynced,
		workqueue:              workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "NamespaceRequests"),
		recorder:               recorder,
	}

	logger.Info("Setting up event handlers")
	// Set up an event handler for when Foo resources change
	namespaceRequestInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueNamespaceRequest,
		UpdateFunc: func(old, new interface{}) {
			controller.enqueueNamespaceRequest(new)
		},
	})
	// Set up an event handler for when Deployment resources change. This
	// handler will lookup the owner of the given Deployment, and if it is
	// owned by a Foo resource then the handler will enqueue that Foo resource for
	// processing. This way, we don't need to implement custom logic for
	// handling Deployment resources. More info on this pattern:
	// https://github.com/kubernetes/community/blob/8cafef897a22026d42f5e5bb3f104febe7e29830/contributors/devel/controllers.md
	// namespaceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
	// 	AddFunc: controller.handleObject,
	// 	UpdateFunc: func(old, new interface{}) {
	// 		newDepl := new.(*appsv1.Deployment)
	// 		oldDepl := old.(*appsv1.Deployment)
	// 		if newDepl.ResourceVersion == oldDepl.ResourceVersion {
	// 			// Periodic resync will send update events for all known Deployments.
	// 			// Two different versions of the same Deployment will always have different RVs.
	// 			return
	// 		}
	// 		controller.handleObject(new)
	// 	},
	// 	DeleteFunc: controller.handleObject,
	// })

	return controller
}

// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers. It will block until stopCh
// is closed, at which point it will shutdown the workqueue and wait for
// workers to finish processing their current work items.
func (c *Controller) Run(ctx context.Context, workers int) error {
	defer utilruntime.HandleCrash()
	defer c.workqueue.ShutDown()
	logger := klog.FromContext(ctx)

	// Start the informer factories to begin populating the informer caches
	logger.Info("Starting Foo controller")

	// Wait for the caches to be synced before starting workers
	logger.Info("Waiting for informer caches to sync")

	if ok := cache.WaitForCacheSync(ctx.Done(), c.namespaceRequestSynced, c.namespaceRequestSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	logger.Info("Starting workers", "count", workers)
	// Launch two workers to process Foo resources
	for i := 0; i < workers; i++ {
		go wait.UntilWithContext(ctx, c.runWorker, time.Second)
	}

	logger.Info("Started workers")
	<-ctx.Done()
	logger.Info("Shutting down workers")

	return nil
}

// runWorker is a long-running function that will continually call the
// processNextWorkItem function in order to read and process a message on the
// workqueue.
func (c *Controller) runWorker(ctx context.Context) {
	for c.processNextWorkItem(ctx) {
	}
}

// processNextWorkItem will read a single work item off the workqueue and
// attempt to process it, by calling the syncHandler.
func (c *Controller) processNextWorkItem(ctx context.Context) bool {
	obj, shutdown := c.workqueue.Get()
	logger := klog.FromContext(ctx)

	if shutdown {
		return false
	}

	// We wrap this block in a func so we can defer c.workqueue.Done.
	err := func(obj interface{}) error {
		// We call Done here so the workqueue knows we have finished
		// processing this item. We also must remember to call Forget if we
		// do not want this work item being re-queued. For example, we do
		// not call Forget if a transient error occurs, instead the item is
		// put back on the workqueue and attempted again after a back-off
		// period.
		defer c.workqueue.Done(obj)
		var key string
		var ok bool
		// We expect strings to come off the workqueue. These are of the
		// form namespace/name. We do this as the delayed nature of the
		// workqueue means the items in the informer cache may actually be
		// more up to date that when the item was initially put onto the
		// workqueue.
		if key, ok = obj.(string); !ok {
			// As the item in the workqueue is actually invalid, we call
			// Forget here else we'd go into a loop of attempting to
			// process a work item that is invalid.
			c.workqueue.Forget(obj)
			utilruntime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}
		// Run the syncHandler, passing it the namespace/name string of the
		// Foo resource to be synced.
		if err := c.syncHandler(ctx, key); err != nil {
			// Put the item back on the workqueue to handle any transient errors.
			c.workqueue.AddRateLimited(key)
			return fmt.Errorf("error syncing '%s': %s, requeuing", key, err.Error())
		}
		// Finally, if no error occurs we Forget this item so it does not
		// get queued again until another change happens.
		c.workqueue.Forget(obj)
		logger.Info("Successfully synced", "resourceName", key)
		return nil
	}(obj)

	if err != nil {
		utilruntime.HandleError(err)
		return true
	}

	return true
}

// syncHandler compares the actual state with the desired, and attempts to
// converge the two. It then updates the Status block of the Foo resource
// with the current status of the resource.
func (c *Controller) syncHandler(ctx context.Context, key string) error {
	// Convert the namespace/name string into a distinct namespace and name
	// logger := klog.LoggerWithValues(klog.FromContext(ctx), "resourceName", key)

	_, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	namespaceRequest, err := c.namespaceRequestLister.Get(name)

	if err != nil {
		if errors.IsNotFound(err) {
			utilruntime.HandleError(fmt.Errorf("namespacerequest '%s' in work queue no longer exists", key))
			return nil
		}

		return err
	}

	c.processNamespaceRequest(namespaceRequest.DeepCopy())
	c.recorder.Event(namespaceRequest, corev1.EventTypeNormal, successSynced, messageResourceSynced)
	return nil
}

func (c *Controller) enqueueNamespaceRequest(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		utilruntime.HandleError(err)
		return
	}
	c.workqueue.Add(key)
}

func (c *Controller) processNamespaceRequest(namespaceRequestCopy *v1alpha1.NamespaceRequest) {
	if namespaceRequestCopy.Status.State == "" {
		namespaceRequestCopy.Status.State = v1alpha1.NamespaceStatusPending
		namespaceRequestCopy.Status.Message = "Waiting for approval"

		c.updateStatus(context.TODO(), namespaceRequestCopy)
		return
	}

	switch namespaceRequestCopy.Status.State {
	case v1alpha1.NamespaceStatusPending:
		// If it is not approved then do nothing
		if namespaceRequestCopy.Spec.Approved {
			if c.validateNamespaceName(context.TODO(), namespaceRequestCopy.Spec.NamespaceName) {
				c.createNamespace(context.TODO(), namespaceRequestCopy)

				namespaceRequestCopy.Status.State = v1alpha1.NamespaceStatusCreated
				namespaceRequestCopy.Status.Message = "Given namespace is created"

				c.updateStatus(context.TODO(), namespaceRequestCopy)
				return
			} else {
				namespaceRequestCopy.Status.State = v1alpha1.NamespaceStatusError
				namespaceRequestCopy.Status.Message = "Given namespace name is not valid or already exists"

				c.updateStatus(context.TODO(), namespaceRequestCopy)
				return
			}
		}
	// Record the event if it is created
	case v1alpha1.NamespaceStatusCreated:
		c.recorder.Event(namespaceRequestCopy, corev1.EventTypeNormal, string(v1alpha1.NamespaceStatusCreated), "Given namespace is created")
	}

}

// Check if the namespace doesn't exist and name is valid
func (c *Controller) validateNamespaceName(ctx context.Context, namespacename string) bool {
	if namespacename == "" {
		return false
	}

	if _, err := c.kubeclientset.CoreV1().Namespaces().Get(ctx, namespacename, v1.GetOptions{}); err == nil {
		return false
	}

	return true
}

func (c *Controller) createNamespace(ctx context.Context, namespaceRequestCopy *v1alpha1.NamespaceRequest) {
	ns := &corev1.Namespace{
		ObjectMeta: v1.ObjectMeta{
			Name: namespaceRequestCopy.Spec.NamespaceName,
		},
	}
	if _, err := c.kubeclientset.CoreV1().Namespaces().Create(ctx, ns, v1.CreateOptions{}); err != nil {
		klog.Infoln(err)
	}
}

// updateStatus calls the API to update the cluster role request status.
func (c *Controller) updateStatus(ctx context.Context, namespaceRequestCopy *v1alpha1.NamespaceRequest) {
	if _, err := c.ubombarclientset.NamespacerequestV1alpha1().NamespaceRequests().UpdateStatus(ctx, namespaceRequestCopy, v1.UpdateOptions{}); err != nil {
		klog.Infoln(err)
	}
}
