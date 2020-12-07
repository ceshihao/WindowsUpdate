package windowsupdate

import (
	"time"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// IUpdate contains the properties and methods that are available to an update.
// https://docs.microsoft.com/zh-cn/windows/win32/api/wuapi/nn-wuapi-iupdate
type IUpdate struct {
	disp                            *ole.IDispatch
	AutoSelectOnWebSites            bool
	BundledUpdates                  []string
	CanRequireSource                bool
	Categories                      []*ICategory
	Deadline                        *time.Time
	DeltaCompressedContentAvailable bool
	DeltaCompressedContentPreferred bool
	DeploymentAction                int32 // enum https://docs.microsoft.com/zh-cn/windows/win32/api/wuapi/ne-wuapi-deploymentaction
	Description                     string
	DownloadContents                []*IUpdateDownloadContent
	DownloadPriority                int32 // enum https://docs.microsoft.com/zh-cn/windows/win32/api/wuapi/ne-wuapi-downloadpriority
	EulaAccepted                    bool
	EulaText                        string
	HandlerID                       string
	Identity                        *IUpdateIdentity
	Image                           *IImageInformation
	InstallationBehavior            *IInstallationBehavior
	IsBeta                          bool
	IsDownloaded                    bool
	IsHidden                        bool
	IsInstalled                     bool
	IsMandatory                     bool
	IsUninstallable                 bool
	KBArticleIDs                    []string
	Languages                       []string
	LastDeploymentChangeTime        *time.Time
	MaxDownloadSize                 int64
	MinDownloadSize                 int64
	MoreInfoUrls                    []string
	MsrcSeverity                    string
	RecommendedCpuSpeed             int32
	RecommendedHardDiskSpace        int32
	RecommendedMemory               int32
	ReleaseNotes                    string
	SecurityBulletinIDs             []string
	SupersededUpdateIDs             []string
	SupportUrl                      string
	Title                           string
	UninstallationBehavior          *IInstallationBehavior
	UninstallationNotes             string
	UninstallationSteps             []string
}

func toIUpdates(updatesDisp *ole.IDispatch) ([]*IUpdate, error) {
	count, err := toInt32Err(oleutil.GetProperty(updatesDisp, "Count"))
	if err != nil {
		return nil, err
	}

	updates := make([]*IUpdate, 0, count)
	for i := 0; i < int(count); i++ {
		updateDisp, err := toIDispatchErr(oleutil.GetProperty(updatesDisp, "Item", i))
		if err != nil {
			return nil, err
		}

		update, err := toIUpdate(updateDisp)
		if err != nil {
			return nil, err
		}

		updates = append(updates, update)
	}
	return updates, nil
}

func toIUpdate(updateDisp *ole.IDispatch) (*IUpdate, error) {
	var err error
	iUpdate := &IUpdate{
		disp: updateDisp,
	}

	if iUpdate.AutoSelectOnWebSites, err = toBoolErr(oleutil.GetProperty(updateDisp, "AutoSelectOnWebSites")); err != nil {
		return nil, err
	}

	if iUpdate.BundledUpdates, err = toStringSliceErr(oleutil.GetProperty(updateDisp, "BundledUpdates")); err != nil {
		return nil, err
	}

	if iUpdate.CanRequireSource, err = toBoolErr(oleutil.GetProperty(updateDisp, "CanRequireSource")); err != nil {
		return nil, err
	}

	categoriesDisp, err := toIDispatchErr(oleutil.GetProperty(updateDisp, "Categories"))
	if err != nil {
		return nil, err
	}
	if iUpdate.Categories, err = toICategories(categoriesDisp); err != nil {
		return nil, err
	}

	if iUpdate.Deadline, err = toTimeErr(oleutil.GetProperty(updateDisp, "Deadline")); err != nil {
		return nil, err
	}

	if iUpdate.DeltaCompressedContentAvailable, err = toBoolErr(oleutil.GetProperty(updateDisp, "DeltaCompressedContentAvailable")); err != nil {
		return nil, err
	}

	if iUpdate.DeltaCompressedContentPreferred, err = toBoolErr(oleutil.GetProperty(updateDisp, "DeltaCompressedContentPreferred")); err != nil {
		return nil, err
	}

	if iUpdate.DeploymentAction, err = toInt32Err(oleutil.GetProperty(updateDisp, "DeploymentAction")); err != nil {
		return nil, err
	}

	if iUpdate.Description, err = toStringErr(oleutil.GetProperty(updateDisp, "Description")); err != nil {
		return nil, err
	}

	downloadContentsDisp, err := toIDispatchErr(oleutil.GetProperty(updateDisp, "DownloadContents"))
	if err != nil {
		return nil, err
	}
	if iUpdate.DownloadContents, err = toIUpdateDownloadContents(downloadContentsDisp); err != nil {
		return nil, err
	}

	if iUpdate.DownloadPriority, err = toInt32Err(oleutil.GetProperty(updateDisp, "DownloadPriority")); err != nil {
		return nil, err
	}

	if iUpdate.EulaAccepted, err = toBoolErr(oleutil.GetProperty(updateDisp, "EulaAccepted")); err != nil {
		return nil, err
	}

	if iUpdate.EulaText, err = toStringErr(oleutil.GetProperty(updateDisp, "EulaText")); err != nil {
		return nil, err
	}

	if iUpdate.HandlerID, err = toStringErr(oleutil.GetProperty(updateDisp, "HandlerID")); err != nil {
		return nil, err
	}

	identityDisp, err := toIDispatchErr(oleutil.GetProperty(updateDisp, "Identity"))
	if err != nil {
		return nil, err
	}
	if iUpdate.Identity, err = toIUpdateIdentity(identityDisp); err != nil {
		return nil, err
	}

	imageDisp, err := toIDispatchErr(oleutil.GetProperty(updateDisp, "Image"))
	if err != nil {
		return nil, err
	}
	if iUpdate.Image, err = toIImageInformation(imageDisp); err != nil {
		return nil, err
	}

	installationBehaviorDisp, err := toIDispatchErr(oleutil.GetProperty(updateDisp, "InstallationBehavior"))
	if err != nil {
		return nil, err
	}
	if iUpdate.InstallationBehavior, err = toIInstallationBehavior(installationBehaviorDisp); err != nil {
		return nil, err
	}

	if iUpdate.IsBeta, err = toBoolErr(oleutil.GetProperty(updateDisp, "IsBeta")); err != nil {
		return nil, err
	}

	if iUpdate.IsDownloaded, err = toBoolErr(oleutil.GetProperty(updateDisp, "IsDownloaded")); err != nil {
		return nil, err
	}

	if iUpdate.IsHidden, err = toBoolErr(oleutil.GetProperty(updateDisp, "IsHidden")); err != nil {
		return nil, err
	}

	if iUpdate.IsInstalled, err = toBoolErr(oleutil.GetProperty(updateDisp, "IsInstalled")); err != nil {
		return nil, err
	}

	if iUpdate.IsMandatory, err = toBoolErr(oleutil.GetProperty(updateDisp, "IsMandatory")); err != nil {
		return nil, err
	}

	if iUpdate.IsUninstallable, err = toBoolErr(oleutil.GetProperty(updateDisp, "IsUninstallable")); err != nil {
		return nil, err
	}

	if iUpdate.KBArticleIDs, err = toStringSliceErr(oleutil.GetProperty(updateDisp, "KBArticleIDs")); err != nil {
		return nil, err
	}

	if iUpdate.Languages, err = toStringSliceErr(oleutil.GetProperty(updateDisp, "Languages")); err != nil {
		return nil, err
	}

	if iUpdate.LastDeploymentChangeTime, err = toTimeErr(oleutil.GetProperty(updateDisp, "LastDeploymentChangeTime")); err != nil {
		return nil, err
	}

	if iUpdate.MaxDownloadSize, err = toInt64Err(oleutil.GetProperty(updateDisp, "MaxDownloadSize")); err != nil {
		return nil, err
	}

	if iUpdate.MinDownloadSize, err = toInt64Err(oleutil.GetProperty(updateDisp, "MinDownloadSize")); err != nil {
		return nil, err
	}

	if iUpdate.MoreInfoUrls, err = toStringSliceErr(oleutil.GetProperty(updateDisp, "MoreInfoUrls")); err != nil {
		return nil, err
	}

	if iUpdate.MsrcSeverity, err = toStringErr(oleutil.GetProperty(updateDisp, "MsrcSeverity")); err != nil {
		return nil, err
	}

	if iUpdate.RecommendedCpuSpeed, err = toInt32Err(oleutil.GetProperty(updateDisp, "RecommendedCpuSpeed")); err != nil {
		return nil, err
	}

	if iUpdate.RecommendedHardDiskSpace, err = toInt32Err(oleutil.GetProperty(updateDisp, "RecommendedHardDiskSpace")); err != nil {
		return nil, err
	}

	if iUpdate.RecommendedMemory, err = toInt32Err(oleutil.GetProperty(updateDisp, "RecommendedMemory")); err != nil {
		return nil, err
	}

	if iUpdate.ReleaseNotes, err = toStringErr(oleutil.GetProperty(updateDisp, "ReleaseNotes")); err != nil {
		return nil, err
	}

	if iUpdate.SecurityBulletinIDs, err = toStringSliceErr(oleutil.GetProperty(updateDisp, "SecurityBulletinIDs")); err != nil {
		return nil, err
	}

	if iUpdate.SupersededUpdateIDs, err = toStringSliceErr(oleutil.GetProperty(updateDisp, "SupersededUpdateIDs")); err != nil {
		return nil, err
	}

	if iUpdate.SupportUrl, err = toStringErr(oleutil.GetProperty(updateDisp, "SupportUrl")); err != nil {
		return nil, err
	}

	if iUpdate.Title, err = toStringErr(oleutil.GetProperty(updateDisp, "Title")); err != nil {
		return nil, err
	}

	uninstallationBehaviorDisp, err := toIDispatchErr(oleutil.GetProperty(updateDisp, "UninstallationBehavior"))
	if err != nil {
		return nil, err
	}
	if iUpdate.UninstallationBehavior, err = toIInstallationBehavior(uninstallationBehaviorDisp); err != nil {
		return nil, err
	}

	if iUpdate.UninstallationNotes, err = toStringErr(oleutil.GetProperty(updateDisp, "UninstallationNotes")); err != nil {
		return nil, err
	}

	if iUpdate.UninstallationSteps, err = toStringSliceErr(oleutil.GetProperty(updateDisp, "UninstallationSteps")); err != nil {
		return nil, err
	}

	return iUpdate, nil
}
