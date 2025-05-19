# vcfReader:
---
CLI tool intended to convert the AGD35k Mitochondrial chromosome sequencing to the formatted expected by Haplogrep3. This tool was implemented in go for better speed and efficiency.

## Compiling code:
---
To compile the code you can just switch into the parent directory and then runt he following command:

```bash
go build .
```

Go has the ability to compile cross platform applications. If you need to compile for an OS that is different than the one you are developing on then you can pass the appropriate environmental variables.

## Example Command:
---
Once compiled you can run the following command to recode the data:

```bash
./vcfRecoder input_file_name.gz output_file_name.gz
```

Replace "input_file_name.gz" and "output_file_name.gz" with your filepaths. Both files are expected to be either gzipped of bgzipped.

