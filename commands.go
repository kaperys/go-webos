package webos

import "github.com/mitchellh/mapstructure"

// Command is the type used by tv.Command to interact with the TV.
type Command string

// APIServiceListCommand lists the API services available on the TV.
const APIServiceListCommand Command = "ssap://api/getServiceList"

// ApplicationManagerForegroundAppCommand returns information about the forgeground app.
const ApplicationManagerForegroundAppCommand Command = "ssap://com.webos.applicationManager/getForegroundAppInfo"

// AudioGetVolumeCommand returns information about the TV's configured audio output volume.
const AudioGetVolumeCommand Command = "ssap://audio/getVolume"

// AudioSetVolumeCommand sets the TV's configured audio output volume.
const AudioSetVolumeCommand Command = "ssap://audio/setVolume"

// AudioVolumeDownCommand decrements the TV's configured audio output volume.
const AudioVolumeDownCommand Command = "ssap://audio/volumeDown"

// AudioVolumeStatusCommand returns information about the TV's configured audio output volume.
// Same as AudioGetVolumeCommand.
const AudioVolumeStatusCommand Command = "ssap://audio/getVolume"

// AudioVolumeUpCommand increments the TV's configured audio output volume.
const AudioVolumeUpCommand Command = "ssap://audio/volumeUp"

// AudioVolumeSetMuteCommand sets/toggles muting the TV's configured audio output.
const AudioVolumeSetMuteCommand Command = "ssap://audio/setMute"

// MediaControlFastForwardCommand fast forwards the current media.
const MediaControlFastForwardCommand Command = "ssap://media.controls/fastForward"

// MediaControlPauseCommand pauses the current media.
const MediaControlPauseCommand Command = "ssap://media.controls/pause"

// MediaControlPlayCommand plays or resumes the current media.
const MediaControlPlayCommand Command = "ssap://media.controls/play"

// MediaControlRewindCommand rewinds the current media.
const MediaControlRewindCommand Command = "ssap://media.controls/rewind"

// MediaControlStopCommand stops the current media.
const MediaControlStopCommand Command = "ssap://media.controls/stop"

// SystemLauncherCloseCommand closes a given application.
const SystemLauncherCloseCommand Command = "ssap://system.launcher/close"

// SystemLauncherGetAppStateCommand returns information about the given application state.
const SystemLauncherGetAppStateCommand Command = "ssap://system.launcher/getAppState"

// SystemLauncherLaunchCommand launches the given application.
const SystemLauncherLaunchCommand Command = "ssap://system.launcher/launch"

// SystemLauncherOpenCommand opens a previously launched application.
const SystemLauncherOpenCommand Command = "ssap://system.launcher/open"

// SystemNotificationsCreateToastCommand creates a "toast" notification.
const SystemNotificationsCreateToastCommand Command = "ssap://system.notifications/createToast"

// SystemTurnOffCommand turns the TV off.
const SystemTurnOffCommand Command = "ssap://system/turnOff"

// TVChannelDownCommand changes the channel down.
const TVChannelDownCommand Command = "ssap://tv/channelDown"

// TVChannelListCommand returns information about the available channels.
const TVChannelListCommand Command = "ssap://tv/getChannelList"

// TVChannelUpCommand changes the channel up.
const TVChannelUpCommand Command = "ssap://tv/channelUp"

// TVCurrentChannelCommand returns information about the current channel.
const TVCurrentChannelCommand Command = "ssap://tv/getCurrentChannel"

// TVCurrentChannelProgramCommand returns information about the current program playing on
// the current channel.
const TVCurrentChannelProgramCommand Command = "ssap://tv/getChannelProgramInfo"

// ServiceList returns information about the available services.
func (tv *TV) ServiceList() (*ServiceList, error) {
	msg, err := tv.Command(APIServiceListCommand, nil)
	if err != nil {
		return nil, err
	}

	sl := &ServiceList{}
	err = mapstructure.Decode(msg.Payload, sl)
	return sl, err
}

// CurrentApp returns information about the current app.
func (tv *TV) CurrentApp() (*App, error) {
	msg, err := tv.Command(ApplicationManagerForegroundAppCommand, nil)
	if err != nil {
		return nil, err
	}

	a := &App{}
	err = mapstructure.Decode(msg.Payload, a)
	return a, err
}

// GetVolume returns information about the audio output volume.
func (tv *TV) GetVolume() (*Volume, error) {
	msg, err := tv.Command(AudioGetVolumeCommand, nil)
	if err != nil {
		return nil, err
	}

	v := &Volume{}
	err = mapstructure.Decode(msg.Payload, v)
	return v, err
}

// SetVolume sets the audio output volume to v.
func (tv *TV) SetVolume(v int) error {
	_, err := tv.Command(AudioSetVolumeCommand, Payload{"volume": v})
	return err
}

// VolumeDown decrements the audio output volume.
func (tv *TV) VolumeDown() error {
	_, err := tv.Command(AudioVolumeDownCommand, nil)
	return err
}

// VolumeStatus returns information about the audio output volume.
func (tv *TV) VolumeStatus() (*Volume, error) {
	msg, err := tv.Command(AudioVolumeStatusCommand, nil)
	if err != nil {
		return nil, err
	}

	v := &Volume{}
	err = mapstructure.Decode(msg.Payload, v)
	return v, err
}

// VolumeUp increments the audio output volume.
func (tv *TV) VolumeUp() error {
	_, err := tv.Command(AudioVolumeUpCommand, nil)
	return err
}

// Mute mutes the TV audio output.
func (tv *TV) Mute() error {
	_, err := tv.Command(AudioVolumeSetMuteCommand, Payload{"mute": 1})
	return err
}

// Unmute unmutes the TV audio output.
func (tv *TV) Unmute() error {
	_, err := tv.Command(AudioVolumeSetMuteCommand, Payload{"mute": 0})
	return err
}

// FastForward fast forwards the current media.
func (tv *TV) FastForward() error {
	_, err := tv.Command(MediaControlFastForwardCommand, nil)
	return err
}

// Pause pauses the current media.
func (tv *TV) Pause() error {
	_, err := tv.Command(MediaControlPauseCommand, nil)
	return err
}

// Play plays or resumes the current media.
func (tv *TV) Play() error {
	_, err := tv.Command(MediaControlPlayCommand, nil)
	return err
}

// Rewind rewinds the current media.
func (tv *TV) Rewind() error {
	_, err := tv.Command(MediaControlRewindCommand, nil)
	return err
}

// Stop stops the current media.
func (tv *TV) Stop() error {
	_, err := tv.Command(MediaControlStopCommand, nil)
	return err
}

// CloseApp closes the given app.
func (tv *TV) CloseApp(app string) error {
	_, err := tv.Command(SystemLauncherCloseCommand, Payload{"id": app})
	return err
}

// AppStatus returns information about the given app status.
func (tv *TV) AppStatus(app string) (*App, error) {
	msg, err := tv.Command(SystemLauncherGetAppStateCommand, Payload{"id": app})
	if err != nil {
		return nil, err
	}

	a := &App{}
	err = mapstructure.Decode(msg.Payload, a)
	return a, err
}

// LaunchApp launches an app.
func (tv *TV) LaunchApp(app string) error {
	_, err := tv.Command(SystemLauncherLaunchCommand, Payload{"id": app})
	return err
}

// OpenApp switches to a previously launched/backgrounded app.
func (tv *TV) OpenApp(app string) error {
	_, err := tv.Command(SystemLauncherOpenCommand, Payload{"id": app})
	return err
}

// Notification creates a "toast" notification.
func (tv *TV) Notification(m string) error {
	_, err := tv.Command(SystemNotificationsCreateToastCommand, Payload{"message": m})
	return err
}

// Shutdown turns the TV off.
func (tv *TV) Shutdown() error {
	_, err := tv.Command(SystemTurnOffCommand, nil)
	return err
}

// ChannelDown decrements the current channel.
func (tv *TV) ChannelDown() error {
	_, err := tv.Command(TVChannelDownCommand, nil)
	return err
}

// ChannelList returns information about available channels.
//  @todo implement a ChannelList type. This doesn't work on my TV.
func (tv *TV) ChannelList() (Message, error) {
	return tv.Command(TVChannelListCommand, nil)
}

// ChannelUp increments the current channel.
func (tv *TV) ChannelUp() error {
	_, err := tv.Command(TVChannelUpCommand, nil)
	return err
}

// CurrentChannel returns information about the current channel.
//  @todo implement a Channel type. This doesn't work on my TV.
func (tv *TV) CurrentChannel() (Message, error) {
	return tv.Command(TVCurrentChannelCommand, nil)
}

// CurrentProgram returns information about the current program
// shown on the CurrentChannel.
//  @todo implement a Program type. This doesn't work on my TV.
func (tv *TV) CurrentProgram() (Message, error) {
	return tv.Command(TVCurrentChannelProgramCommand, nil)
}
