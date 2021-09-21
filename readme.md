Dumps the Downstream and Upstream channel details for the CM500-100NAS modem

## Building

`./build.sh`

## Running

`./modemstat`

## Demo
```
$ ./modemstat

Downstream Bonded Channels
==========================
 4  Locked  QAM256  501000000 Hz  -0.1 dBmV  35.6 dB         6310          283
 1  Locked  QAM256  483000000 Hz  -0.4 dBmV  47.4 dB            0            0
 2  Locked  QAM256  489000000 Hz  -0.3 dBmV  47.4 dB            0            0
 3  Locked  QAM256  495000000 Hz  -0.3 dBmV  46.9 dB            0            0
 5  Locked  QAM256  507000000 Hz  -0.2 dBmV  39.7 dB            0            0
 6  Locked  QAM256  513000000 Hz  -0.1 dBmV  38.5 dB            0            0
 7  Locked  QAM256  519000000 Hz     0 dBmV  33.1 dB     39664639      4285696
 8  Locked  QAM256  525000000 Hz     0 dBmV  30.1 dB     10683712      2563767
 9  Locked  QAM256  531000000 Hz   0.1 dBmV  29.8 dB     20291209      2979755
10  Locked  QAM256  537000000 Hz     0 dBmV  35.8 dB       364143         3779
11  Locked  QAM256  543000000 Hz     0 dBmV  39.5 dB            0            0
12  Locked  QAM256  549000000 Hz     0 dBmV  44.1 dB            0            0
13  Locked  QAM256  555000000 Hz  -0.1 dBmV  46.9 dB            0            0
14  Locked  QAM256  561000000 Hz  -0.1 dBmV  47.1 dB            0            0
15  Locked  QAM256  567000000 Hz  -0.1 dBmV  45.4 dB            0            0
16  Locked  QAM256  573000000 Hz     0 dBmV  44.1 dB            0            0

Upstream Bonded Channels
==========================
 4  Locked  ATDMA  5120 Ksym/sec  16600000 Hz  41.1 dBmV
 2  Locked  ATDMA  5120 Ksym/sec  29400000 Hz  42.2 dBmV
 3  Locked  ATDMA  5120 Ksym/sec  23000000 Hz  41.5 dBmV
 1  Locked  ATDMA  5120 Ksym/sec  35800000 Hz  42.7 dBmV
```