# musicTracker

a simple websocket server which keeps track of currently playing music track. clients can update the track and monitors can request track info. made to host information for an esp32 board with a screen to visualize tracks.

> [!IMPORTANT]
> a .env file should be created with these lines:
```
ip={your_ip}
port={your_port}
```

- on connect you should set your id with `!id:{id}`, which can be **Monitor** or **Client**
- to get your own id: `?id`
- to update track info: `!track:{title}|{artist}|{album}|{yt_video_id}|{playtime_in_seconds}`
- to get track info: `?track`
- to get the client's current progress in the track: ?time
- to send the client's current progress in the track: !time:{time_in_seconds}

the **{yt_video_id}** is needed for easy access to the track's thumbnail
