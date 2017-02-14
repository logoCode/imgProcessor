# imgProcessor

## description

This command line application is intended to be used to display the output of Simplex3 experiments visually on coordinate systems. Although starting positions of some object are meant to be aligned to the axis of the coordinate system, any other value can be used.  So, while the axis describe input values of the model, the color of each pixel indicates its final state. 

## documentation

To get the output of a Simplex3 experiment, you have to use the console. With the DISPLAY() method of simplex-edl, you can print the important information to the console. Afterwards you have to copy the whole output to some file, to hand it over to imgProcessor. To label the lines which include relevant output, you have to begin each line with some identifier like e.g. "$Data". Each output line has to consist of three parameters, the two input values, and the output value, which describes the final state of the model. The idenfifier and each of the three parameters have to be seperated with a separator like "/". An Example could look like this:
```
DISPLAY("$Data / %lg / %lg / %d \n", x, y, endVal);
```
The identifier and the separator can be changed in the settings of imgProcessor. There you can also specify the amount of decimals the input values have as a maximum, so that the output image is scaled correctly, and all the information is used. Even though there are seven default colors, you can specify as many as you want.

After you enter the command "process" in imgProcessor, the image will be created and saved. The output name and location can be specified in the settings.
