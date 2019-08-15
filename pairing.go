package webos

// pairPrompt returns a Payload necessary to pair with the TV using
// the PROMPT method.
func pairPrompt() Payload {
	return Payload{
		"forcePairing": false,
		"pairingType":  "PROMPT",
		"manifest": map[string]interface{}{
			"manifestVersion": 1,
			"appVersion":      "1.1",
			"signed": map[string]interface{}{
				"created":  "20140509",
				"appId":    "com.lge.test",
				"vendorId": "com.lge",
				"localizedAppNames": map[string]string{
					"": "LG Remote App",
				},
				"localizedVendorNames": map[string]string{
					"": "LG Electronics",
				},
				"permissions": []string{
					"LAUNCH",
					"LAUNCH_WEBAPP",
					"APP_TO_APP",
					"CLOSE",
					"TEST_OPEN",
					"TEST_PROTECTED",
					"CONTROL_AUDIO",
					"CONTROL_DISPLAY",
					"CONTROL_INPUT_JOYSTICK",
					"CONTROL_INPUT_MEDIA_RECORDING",
					"CONTROL_INPUT_MEDIA_PLAYBACK",
					"CONTROL_INPUT_TV",
					"CONTROL_POWER",
					"READ_APP_STATUS",
					"READ_CURRENT_CHANNEL",
					"READ_INPUT_DEVICE_LIST",
					"READ_NETWORK_STATE",
					"READ_RUNNING_APPS",
					"READ_TV_CHANNEL_LIST",
					"WRITE_NOTIFICATION_TOAST",
					"READ_POWER_STATE",
					"READ_COUNTRY_INFO",
				},
				"serial": "2f930e2d2cfe083771f68e4fe7bb07",
			},
			"permissions": []string{
				"TEST_SECURE",
				"CONTROL_INPUT_TEXT",
				"CONTROL_MOUSE_AND_KEYBOARD",
				"READ_INSTALLED_APPS",
				"READ_LGE_SDX",
				"READ_NOTIFICATIONS",
				"SEARCH",
				"WRITE_SETTINGS",
				"WRITE_NOTIFICATION_ALERT",
				"CONTROL_POWER",
				"READ_CURRENT_CHANNEL",
				"READ_RUNNING_APPS",
				"READ_UPDATE_INFO",
				"UPDATE_FROM_REMOTE_APP",
				"READ_LGE_TV_INPUT_EVENTS",
				"READ_TV_CURRENT_TIME",
			},
			"signatures": []map[string]interface{}{
				{
					"signatureVersion": 1,
					"signature":        "eyJhbGdvcml0aG0iOiJSU0EtU0hBMjU2Iiwia2V5SWQiOiJ0ZXN0LXNpZ25pbmctY2VydCIsInNpZ25hdHVyZVZlcnNpb24iOjF9.hrVRgjCwXVvE2OOSpDZ58hR+59aFNwYDyjQgKk3auukd7pcegmE2CzPCa0bJ0ZsRAcKkCTJrWo5iDzNhMBWRyaMOv5zWSrthlf7G128qvIlpMT0YNY+n/FaOHE73uLrS/g7swl3/qH/BGFG2Hu4RlL48eb3lLKqTt2xKHdCs6Cd4RMfJPYnzgvI4BNrFUKsjkcu+WD4OO2A27Pq1n50cMchmcaXadJhGrOqH5YmHdOCj5NSHzJYrsW0HPlpuAx/ECMeIZYDh6RMqaFM2DXzdKX9NmmyqzJ3o/0lkk/N97gfVRLW5hA29yeAwaCViZNCP8iC9aO0q9fQojoa7NQnAtw==",
				},
			},
		},
	}
}

// pairPIN returns a Payload necessary to pair with the TV using
// the PIN method.
//  @todo implement PIN pairing.
func pairPIN() Payload {
	panic("not implemented")
}
