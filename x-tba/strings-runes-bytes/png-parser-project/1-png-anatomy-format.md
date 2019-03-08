# The Brief Anatomy of a PNG image

```
The first 24 bytes:

+=================================+
|  PNG Header         |  8 bytes  | -> 89 50 4e 47 0d 0a 1a 0a
+---------------------+-----------+
|  IHDR Chunk Header  |           |
|    Chunk Length     |  4 bytes  | -> The length of the IHDR Chunk Data
|    Chunk Type       |  4 bytes  | -> 49 48 44 52
+---------------------+-----------+
|  IHDR Chunk Data    |           |
|    Width            |  4 bytes  | -> uint32 — big endian
|    Height           |  4 bytes  | -> uint32 — big endian
+=================================+
```