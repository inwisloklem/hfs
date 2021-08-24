package core

import (
	"errors"
	"fmt"
	"hfs/settings"
	"os"
	"time"

	"github.com/otiai10/copy"
)

const (
	storeDirName = "Store"
	savesDirName = "Saves"
)

type Control struct {
	Config      settings.Config
	HasNoConfig bool
	Message     string
}

func (c *Control) CheckDir(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return errors.New(c.GetProfileName() + " dir doesn't exist")
	}
	return nil
}

func (c *Control) GetHasNoConfig() bool {
	return c.HasNoConfig
}

func (c *Control) GetDirPath(dirName string) string {
	return c.Config.Path + "\\" + dirName
}

func (c *Control) GetMessage() string {
	return c.Message
}

func (c *Control) GetProfileName() string {
	return "Profile_" + fmt.Sprint(c.Config.ProfileNumber)
}

func (c *Control) GetProfilePath(dirName string) string {
	return c.GetDirPath(dirName) + "\\" + c.GetProfileName()
}

func (c *Control) SetErrorMessage(err error) {
	c.Message = "Error: " + err.Error()
}

func GetFormattedTime() string {
	t := time.Now()
	return fmt.Sprintf("_%d%02d%02d_%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func (c *Control) GetStoreDirSavesNumber() int {
	entries, err := os.ReadDir(c.GetDirPath(storeDirName))
	if err != nil {
		return 0
	}
	return len(entries)
}

func (c *Control) LoadFile() {
	entries, err := os.ReadDir(c.GetDirPath(storeDirName))
	if err != nil {
		c.SetErrorMessage(errors.New("can't read files in \"" + storeDirName + "\" dir"))
		return
	}
	if len(entries) == 0 {
		c.SetErrorMessage(errors.New("no files in \"" + storeDirName + "\" dir"))
		return
	}
	profilePath := c.GetProfilePath(savesDirName)
	err = os.RemoveAll(profilePath)
	if err != nil {
		c.SetErrorMessage(errors.New("can't remove old " + c.GetProfileName()))
		return
	}
	lastEntry := entries[len(entries)-1]
	err = copy.Copy(c.GetDirPath("Store")+"\\"+lastEntry.Name(), profilePath)
	if err != nil {
		c.SetErrorMessage(errors.New("can't copy from \"" + storeDirName + "\" to \"" + savesDirName + "\" dir "))
		return
	}
	c.Message = lastEntry.Name() + " was loaded to " + c.GetProfileName()
}

func (c *Control) SaveFile() {
	err := c.CheckDir(c.GetProfilePath(savesDirName))
	if err != nil {
		c.SetErrorMessage(err)
		return
	}
	time := GetFormattedTime()
	err = copy.Copy(c.GetProfilePath(savesDirName), c.GetProfilePath(storeDirName)+time)
	if err != nil {
		c.SetErrorMessage(errors.New("can't copy from \"" + savesDirName + "\" to \"" + storeDirName + "\" dir "))
		return
	}
	profileName := c.GetProfileName()
	c.Message = profileName + " was saved to " + profileName + time
}
