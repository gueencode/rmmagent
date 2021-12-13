package agent

import (
	"fmt"

	rmm "rmmagent/shared"

	wapi "github.com/gueencode/go-win64api"
)

func (a *WindowsAgent) GetInstalledSoftware() []rmm.SoftwareList {
	ret := make([]rmm.SoftwareList, 0)

	sw, err := wapi.InstalledSoftwareList()
	if err != nil {
		return ret
	}

	for _, s := range sw {
		t := s.InstallDate
		ret = append(ret, rmm.SoftwareList{
			Name:        s.Name(),
			Version:     s.Version(),
			Publisher:   s.Publisher,
			InstallDate: fmt.Sprintf("%02d-%d-%02d", t.Year(), t.Month(), t.Day()),
			Size:        ByteCountSI(s.EstimatedSize * 1024),
			Source:      s.InstallSource,
			Location:    s.InstallLocation,
			Uninstall:   s.UninstallString,
		})
	}
	return ret
}
