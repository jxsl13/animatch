# animatch

Early stage.

**animatch** is intended to search and match a whole lot of video files and add either a specific AniDB tag like `[anidb-123456]` or rename all of those files to their proper file names in a Plex appropriate manner.

My current use case is to prepare a bunch of files for FileBot and in the long run maybe to replace FileBot, even tho it is a great tool.

The go templating engine is powerful enough to replace a lot of the functionality that is provided by Groovy used in FileBot.


# FileBot renaming script

Initially you may tag your files
```shell
animatch tag path/to/file/or/folder
```

FileBot bash script
```shell
filebot -rename input-dir \
--output output-dir \
--action test \
--order Airdate \
--conflict auto \
--db AniDB \
--q "{(fn =~ /\[anidb-(\d+)\]/)[0][1]}" \
-non-strict \
--lang "English" \
--log all \
--format "{plex}"
```


