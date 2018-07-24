package webos

type Command string

const ApiServiceListCommand Command = "ssap://api/getServiceList"
const ApplicationManagerForegroundAppCommand Command = "ssap://com.webos.applicationManager/getForegroundAppInfo"
const AudioGetVolumeCommand Command = "ssap://audio/getVolume"
const AudioSetVolumeCommand Command = "ssap://audio/setVolume"
const AudioVolumeDownCommand Command = "ssap://audio/volumeDown"
const AudioVolumeStatusCommand Command = "ssap://audio/getVolume"
const AudioVolumeUpCommand Command = "ssap://audio/volumeUp"
const AudioVoumeSetMuteCommand Command = "ssap://audio/setMute"
const MediaControlFastForwardCommand Command = "ssap://media.controls/fastForward"
const MediaControlPauseCommand Command = "ssap://media.controls/pause"
const MediaControlPlayCommand Command = "ssap://media.controls/play"
const MediaControlRewindCommand Command = "ssap://media.controls/rewind"
const MediaControlStopCommand Command = "ssap://media.controls/stop"
const SystemLauncherCloseCommand Command = "ssap://system.launcher/close"
const SystemLauncherGetAppStateCommand Command = "ssap://system.launcher/getAppState"
const SystemLauncherLaunchCommand Command = "ssap://system.launcher/launch"
const SystemLauncherOpenCommand Command = "ssap://system.launcher/open"
const SystemNotificationsCreateToastCommand Command = "ssap://system.notifications/createToast"
const SystemTurnOffCommand Command = "ssap://system/turnOff"
const TVChannelDownCommand Command = "ssap://tv/channelDown"
const TVChannelListCommand Command = "ssap://tv/getChannelList"
const TVChannelUpCommand Command = "ssap://tv/channelUp"
const TVCurrentChannelCommand Command = "ssap://tv/getCurrentChannel"
const TVCurrentChannelProgramCommand Command = "ssap://tv/getChannelProgramInfo"

func (tv *TV) ServiceList() (Message, error) {
	return tv.Command(ApiServiceListCommand, nil)
}

func (tv *TV) CurrentApp() (Message, error) {
	return tv.Command(ApplicationManagerForegroundAppCommand, nil)
}

func (tv *TV) GetVolume() (Message, error) {
	return tv.Command(AudioGetVolumeCommand, nil)
}

func (tv *TV) SetVolume(v int) error {
	_, err := tv.Command(AudioSetVolumeCommand, Payload{"volume": v})
	return err
}

func (tv *TV) VolumeDown() error {
	_, err := tv.Command(AudioVolumeDownCommand, nil)
	return err
}

func (tv *TV) VolumeStatus() (Message, error) {
	return tv.Command(AudioVolumeStatusCommand, nil)
}

func (tv *TV) VolumeUp() error {
	_, err := tv.Command(AudioVolumeUpCommand, nil)
	return err
}

func (tv *TV) Mute() error {
	_, err := tv.Command(AudioVoumeSetMuteCommand, Payload{"mute": 1})
	return err
}

func (tv *TV) Unmute() error {
	_, err := tv.Command(AudioVoumeSetMuteCommand, Payload{"mute": 0})
	return err
}

func (tv *TV) FastForward() error {
	_, err := tv.Command(MediaControlFastForwardCommand, nil)
	return err
}

func (tv *TV) Pause() error {
	_, err := tv.Command(MediaControlPauseCommand, nil)
	return err
}

func (tv *TV) Play() error {
	_, err := tv.Command(MediaControlPlayCommand, nil)
	return err
}

func (tv *TV) Rewind() error {
	_, err := tv.Command(MediaControlRewindCommand, nil)
	return err
}

func (tv *TV) Stop() error {
	_, err := tv.Command(MediaControlStopCommand, nil)
	return err
}

func (tv *TV) CloseApp(app string) error {
	_, err := tv.Command(SystemLauncherCloseCommand, Payload{"id": app})
	return err
}

func (tv *TV) AppStatus(app string) (Message, error) {
	return tv.Command(SystemLauncherGetAppStateCommand, Payload{"id": app})
}

func (tv *TV) LaunchApp(app string) error {
	_, err := tv.Command(SystemLauncherLaunchCommand, Payload{"id": app})
	return err
}

func (tv *TV) OpenApp(app string) error {
	_, err := tv.Command(SystemLauncherOpenCommand, Payload{"id": app})
	return err
}

func (tv *TV) Notification(m string) error {
	_, err := tv.Command(SystemNotificationsCreateToastCommand, Payload{"message": m})
	return err
}

func (tv *TV) Shutdown() error {
	_, err := tv.Command(SystemTurnOffCommand, nil)
	return err
}

func (tv *TV) ChannelDown() error {
	_, err := tv.Command(TVChannelDownCommand, nil)
	return err
}

func (tv *TV) ChannelList() (Message, error) {
	return tv.Command(TVChannelListCommand, nil)
}

func (tv *TV) ChannelUp() error {
	_, err := tv.Command(TVChannelUpCommand, nil)
	return err
}

func (tv *TV) CurrentChannel() (Message, error) {
	return tv.Command(TVCurrentChannelCommand, nil)
}

func (tv *TV) CurrentProgram() (Message, error) {
	return tv.Command(TVCurrentChannelProgramCommand, nil)
}
