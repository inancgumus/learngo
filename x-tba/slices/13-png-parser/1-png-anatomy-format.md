# The Brief Anatomy of a PNG image

```
The first 24 bytes:

+=================================+
|  PNG Header         |  8 bytes  | -> 137 80 78 71 13 10 26 10
+---------------------+-----------+
|  IHDR Chunk Header  |           |
|    Chunk Length     |  4 bytes  | -> The length of the IHDR Chunk Data
|    Chunk Type       |  4 bytes  | -> 73 72 68 82
+---------------------+-----------+
|  IHDR Chunk Data    |           |
|    Width            |  4 bytes  | -> Unsigned 32-bit integer
|    Height           |  4 bytes  | -> Unsigned 32-bit integer
+=================================+
```