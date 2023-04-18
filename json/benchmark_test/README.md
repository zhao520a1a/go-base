## Benchmarks

### 测试数据
- 小[Small](../testdata/small.go) ：400B，11 key，深度 3 层； 
- 中[Medium](../testdata/medium.go)：13KB，300+ key，深度 4层（实际业务数据，其中有大量的嵌套 JSON string)； 
- 大[Large](../testdata/large.json)：635KB，10000+ key，深度 6 层。

这里参考了 [bytedance/sonic](https://github.com/bytedance/sonic) 库中压测方式

### 压测结果

#### 解析小 JSON 字符串
```
BenchmarkDecoder_Generic_StdLib-12                       1000000              6986 ns/op          52.25 MB/s        3474 B/op         91 allocs/op
BenchmarkDecoder_Generic_JsonIter-12                     1000000              7366 ns/op          49.55 MB/s        3748 B/op        112 allocs/op
BenchmarkDecoder_Generic_GoJson-12                       1000000              6961 ns/op          52.43 MB/s        4279 B/op        105 allocs/op
BenchmarkDecoder_Generic_Sonic-12                        1000000              4093 ns/op          89.17 MB/s        4567 B/op         42 allocs/op
BenchmarkDecoder_Generic_Sonic_V1-12                     1000000              4474 ns/op          81.58 MB/s        4363 B/op         71 allocs/op
BenchmarkDecoder_Generic_Sonic_Fast-12                   1000000              3927 ns/op          92.95 MB/s        4171 B/op         41 allocs/op
BenchmarkDecoder_Generic_Sonnet-12                       1000000              4743 ns/op          76.95 MB/s        3330 B/op         80 allocs/op
BenchmarkDecoder_Parallel_Generic_StdLib-12              1000000              2387 ns/op         152.89 MB/s        3474 B/op         91 allocs/op
BenchmarkDecoder_Parallel_Generic_JsonIter-12            1000000              2411 ns/op         151.37 MB/s        3749 B/op        112 allocs/op
BenchmarkDecoder_Parallel_Generic_GoJson-12              1000000              2680 ns/op         136.20 MB/s        4283 B/op        105 allocs/op
BenchmarkDecoder_Parallel_Generic_Sonic-12               1000000              1641 ns/op         222.45 MB/s        4615 B/op         42 allocs/op
BenchmarkDecoder_Parallel_Generic_Sonic_V1-12            1000000              1707 ns/op         213.82 MB/s        4383 B/op         71 allocs/op
BenchmarkDecoder_Parallel_Generic_Sonic_Fast-12          1000000              1688 ns/op         216.22 MB/s        4190 B/op         41 allocs/op
BenchmarkDecoder_Parallel_Generic_Sonnet-12              1000000              2154 ns/op         169.45 MB/s        3331 B/op         80 allocs/op
BenchmarkDecoder_Binding_StdLib-12                       1000000              7369 ns/op          49.53 MB/s        1056 B/op         36 allocs/op
BenchmarkDecoder_Binding_JsonIter-12                     1000000              2281 ns/op         160.01 MB/s         688 B/op         31 allocs/op
BenchmarkDecoder_Binding_GoJson-12                       1000000              1786 ns/op         204.36 MB/s         939 B/op         19 allocs/op
BenchmarkDecoder_Binding_Sonic-12                        1000000              1697 ns/op         215.14 MB/s        1797 B/op          8 allocs/op
BenchmarkDecoder_Binding_Sonic_V1-12                     1000000              1720 ns/op         212.23 MB/s        1464 B/op         14 allocs/op
BenchmarkDecoder_Binding_Sonic_Fast-12                   1000000              1616 ns/op         225.85 MB/s        1403 B/op          7 allocs/op
BenchmarkDecoder_Binding_Sonnet-12                       1000000              4747 ns/op          76.89 MB/s        3330 B/op         80 allocs/op
BenchmarkDecoder_Parallel_Binding_StdLib-12              1000000              1835 ns/op         198.92 MB/s        1055 B/op         36 allocs/op
BenchmarkDecoder_Parallel_Binding_JsonIter-12            1000000               638.5 ns/op       571.62 MB/s         688 B/op         31 allocs/op
BenchmarkDecoder_Parallel_Binding_GoJson-12              1000000               640.7 ns/op       569.68 MB/s         942 B/op         19 allocs/op
BenchmarkDecoder_Parallel_Binding_Sonic-12               1000000               557.6 ns/op       654.64 MB/s        1868 B/op          8 allocs/op
BenchmarkDecoder_Parallel_Binding_Sonic_V1-12            1000000               523.6 ns/op       697.05 MB/s        1496 B/op         14 allocs/op
BenchmarkDecoder_Parallel_Binding_Sonic_Fast-12          1000000               504.4 ns/op       723.67 MB/s        1436 B/op          7 allocs/op
BenchmarkDecoder_Parallel_Binding_Sonnet-12              1000000               718.4 ns/op       508.06 MB/s         751 B/op         17 allocs/op
BenchmarkEncoder_Generic_StdLib-12                       1000000             10142 ns/op          35.99 MB/s        3273 B/op         67 allocs/op
BenchmarkEncoder_Generic_JsonIter-12                     1000000              5373 ns/op          67.93 MB/s        1024 B/op         16 allocs/op
BenchmarkEncoder_Generic_GoJson-12                       1000000              4903 ns/op          74.44 MB/s         961 B/op          2 allocs/op
BenchmarkEncoder_Generic_Sonic-12                        1000000              2405 ns/op         151.76 MB/s         479 B/op          4 allocs/op
BenchmarkEncoder_Generic_Sonic_V1-12                     1000000              2734 ns/op         133.49 MB/s         732 B/op          4 allocs/op
BenchmarkEncoder_Generic_Sonic_Fast-12                   1000000              2329 ns/op         156.72 MB/s         469 B/op          4 allocs/op
BenchmarkEncoder_Parallel_Generic_Sonic-12               1000000               475.6 ns/op       767.50 MB/s         487 B/op          4 allocs/op
BenchmarkEncoder_Parallel_Generic_JsonIter-12            1000000              1248 ns/op         292.50 MB/s        1024 B/op         16 allocs/op
BenchmarkEncoder_Parallel_Generic_GoJson-12              1000000              1636 ns/op         223.16 MB/s         961 B/op          2 allocs/op
BenchmarkEncoder_Parallel_Generic_StdLib-12              1000000              3971 ns/op          91.92 MB/s        3273 B/op         67 allocs/op
BenchmarkEncoder_Parallel_Generic_Sonic_V1-12            1000000              1003 ns/op         363.81 MB/s         760 B/op          4 allocs/op
BenchmarkEncoder_Parallel_Generic_Sonic_Fast-12          1000000               537.9 ns/op       678.51 MB/s         492 B/op          4 allocs/op
BenchmarkEncoder_Binding_StdLib-12                       1000000              2073 ns/op         176.07 MB/s         384 B/op          1 allocs/op
BenchmarkEncoder_Binding_JsonIter-12                     1000000              2209 ns/op         165.26 MB/s         392 B/op          2 allocs/op
BenchmarkEncoder_Binding_GoJson-12                       1000000              1172 ns/op         311.31 MB/s         384 B/op          1 allocs/op
BenchmarkEncoder_Binding_Sonic-12                        1000000               842.2 ns/op       433.41 MB/s         492 B/op          4 allocs/op
BenchmarkEncoder_Binding_Sonic_V1-12                     1000000               956.8 ns/op       381.50 MB/s         757 B/op          4 allocs/op
BenchmarkEncoder_Binding_Sonic_Fast-12                   1000000               814.0 ns/op       448.41 MB/s         490 B/op          4 allocs/op
BenchmarkEncoder_Parallel_Binding_StdLib-12              1000000               452.7 ns/op       806.23 MB/s         384 B/op          1 allocs/op
BenchmarkEncoder_Parallel_Binding_JsonIter-12            1000000               508.1 ns/op       718.34 MB/s         392 B/op          2 allocs/op
BenchmarkEncoder_Parallel_Binding_GoJson-12              1000000               331.9 ns/op      1099.89 MB/s         384 B/op          1 allocs/op
BenchmarkEncoder_Parallel_Binding_Sonic-12               1000000               184.7 ns/op      1976.14 MB/s         488 B/op          4 allocs/op
BenchmarkEncoder_Parallel_Binding_Sonic_V1-12            1000000               238.7 ns/op      1529.08 MB/s         745 B/op          4 allocs/op
BenchmarkEncoder_Parallel_Binding_Sonic_Fast-12          1000000               185.8 ns/op      1964.13 MB/s         496 B/op          4 allocs/op
```

#### 解析中 JSON 字符串
```
BenchmarkDecoder_Generic_StdLib-12                        100000            134048 ns/op          97.26 MB/s       50870 B/op        772 allocs/op
BenchmarkDecoder_Generic_JsonIter-12                      100000             94131 ns/op         138.50 MB/s       55779 B/op       1068 allocs/op
BenchmarkDecoder_Generic_GoJson-12                        100000             91549 ns/op         142.40 MB/s       66358 B/op        973 allocs/op
BenchmarkDecoder_Generic_Sonic-12                         100000             57213 ns/op         227.87 MB/s       62757 B/op        314 allocs/op
BenchmarkDecoder_Generic_Sonic_V1-12                      100000             67120 ns/op         194.23 MB/s       56800 B/op        723 allocs/op
BenchmarkDecoder_Generic_Sonic_Fast-12                    100000             54601 ns/op         238.77 MB/s       49020 B/op        313 allocs/op
BenchmarkDecoder_Generic_Sonnet-12                        100000             67423 ns/op         193.36 MB/s       50222 B/op        739 allocs/op
BenchmarkDecoder_Parallel_Generic_StdLib-12               100000             50068 ns/op         260.38 MB/s       50875 B/op        772 allocs/op
BenchmarkDecoder_Parallel_Generic_JsonIter-12             100000             38097 ns/op         342.20 MB/s       55790 B/op       1068 allocs/op
BenchmarkDecoder_Parallel_Generic_GoJson-12               100000             33294 ns/op         391.57 MB/s       66364 B/op        974 allocs/op
BenchmarkDecoder_Parallel_Generic_Sonic-12                100000             21363 ns/op         610.25 MB/s       63889 B/op        314 allocs/op
BenchmarkDecoder_Parallel_Generic_Sonic_V1-12             100000             24666 ns/op         528.55 MB/s       56823 B/op        723 allocs/op
BenchmarkDecoder_Parallel_Generic_Sonic_Fast-12           100000             19639 ns/op         663.84 MB/s       49063 B/op        313 allocs/op
BenchmarkDecoder_Parallel_Generic_Sonnet-12               100000             33211 ns/op         392.55 MB/s       50233 B/op        739 allocs/op
BenchmarkDecoder_Binding_StdLib-12                        100000            119862 ns/op         108.77 MB/s       10576 B/op        208 allocs/op
BenchmarkDecoder_Binding_JsonIter-12                      100000             37648 ns/op         346.29 MB/s       14672 B/op        385 allocs/op
BenchmarkDecoder_Binding_GoJson-12                        100000             30700 ns/op         424.66 MB/s       22031 B/op         49 allocs/op
BenchmarkDecoder_Binding_Sonic-12                         100000             31562 ns/op         413.05 MB/s       38060 B/op         35 allocs/op
BenchmarkDecoder_Binding_Sonic_V1-12                      100000             32212 ns/op         404.72 MB/s       27327 B/op        137 allocs/op
BenchmarkDecoder_Binding_Sonic_Fast-12                    100000             28587 ns/op         456.05 MB/s       24280 B/op         34 allocs/op
BenchmarkDecoder_Binding_Sonnet-12                        100000             68974 ns/op         189.01 MB/s       50228 B/op        739 allocs/op
BenchmarkDecoder_Parallel_Binding_StdLib-12               100000             27176 ns/op         479.73 MB/s       10575 B/op        208 allocs/op
BenchmarkDecoder_Parallel_Binding_JsonIter-12             100000             11446 ns/op        1138.96 MB/s       14674 B/op        385 allocs/op
BenchmarkDecoder_Parallel_Binding_GoJson-12               100000              9237 ns/op        1411.39 MB/s       22102 B/op         49 allocs/op
BenchmarkDecoder_Parallel_Binding_Sonic-12                100000              9081 ns/op        1435.69 MB/s       39066 B/op         35 allocs/op
BenchmarkDecoder_Parallel_Binding_Sonic_V1-12             100000              9358 ns/op        1393.20 MB/s       27372 B/op        137 allocs/op
BenchmarkDecoder_Parallel_Binding_Sonic_Fast-12           100000              8292 ns/op        1572.23 MB/s       24347 B/op         34 allocs/op
BenchmarkDecoder_Parallel_Binding_Sonnet-12               100000             11304 ns/op        1153.32 MB/s       11656 B/op        130 allocs/op
BenchmarkEncoder_Generic_StdLib-12                        100000            110818 ns/op         117.64 MB/s       44266 B/op        751 allocs/op
BenchmarkEncoder_Generic_JsonIter-12                      100000             44940 ns/op         290.10 MB/s       14347 B/op        115 allocs/op
BenchmarkEncoder_Generic_GoJson-12                        100000             68311 ns/op         190.85 MB/s       23464 B/op         18 allocs/op
BenchmarkEncoder_Generic_Sonic-12                         100000             33923 ns/op         384.31 MB/s       13759 B/op          4 allocs/op
BenchmarkEncoder_Generic_Sonic_Fast-12                    100000             23436 ns/op         556.28 MB/s        9727 B/op          4 allocs/op
BenchmarkEncoder_Parallel_Generic_StdLib-12               100000             40058 ns/op         325.45 MB/s       44288 B/op        751 allocs/op
BenchmarkEncoder_Parallel_Generic_JsonIter-12             100000             11915 ns/op        1094.20 MB/s       14358 B/op        115 allocs/op
BenchmarkEncoder_Parallel_Generic_GoJson-12               100000             20948 ns/op         622.36 MB/s       23445 B/op         18 allocs/op
BenchmarkEncoder_Parallel_Generic_Sonic-12                100000             10335 ns/op        1261.48 MB/s       14056 B/op          4 allocs/op
BenchmarkEncoder_Parallel_Generic_Sonic_Fast-12           100000              5895 ns/op        2211.64 MB/s        9821 B/op          4 allocs/op
BenchmarkEncoder_Binding_StdLib-12                        100000             18100 ns/op         720.29 MB/s        9480 B/op          1 allocs/op
BenchmarkEncoder_Binding_JsonIter-12                      100000             22245 ns/op         586.08 MB/s        9487 B/op          2 allocs/op
BenchmarkEncoder_Binding_GoJson-12                        100000              7886 ns/op        1653.13 MB/s        9482 B/op          1 allocs/op
BenchmarkEncoder_Binding_Sonic-12                         100000              5659 ns/op        2303.64 MB/s       14253 B/op          4 allocs/op
BenchmarkEncoder_Binding_Sonic_Fast-12                    100000              4961 ns/op        2627.89 MB/s       10056 B/op          4 allocs/op
BenchmarkEncoder_Parallel_Binding_StdLib-12               100000              4468 ns/op        2917.64 MB/s        9487 B/op          1 allocs/op
BenchmarkEncoder_Parallel_Binding_JsonIter-12             100000              5370 ns/op        2427.56 MB/s        9496 B/op          2 allocs/op
BenchmarkEncoder_Parallel_Binding_GoJson-12               100000              2928 ns/op        4453.05 MB/s        9500 B/op          1 allocs/op
BenchmarkEncoder_Parallel_Binding_Sonic-12                100000              1348 ns/op        9670.27 MB/s       14279 B/op          4 allocs/op
BenchmarkEncoder_Parallel_Binding_Sonic_Fast-12           100000              1162 ns/op        11219.60 MB/s      10056 B/op          4 allocs/op
```

#### 解析大 JSON 字符串
```
BenchmarkDecoder_Generic_StdLib-12                          5000           6120476 ns/op         103.18 MB/s     2151910 B/op      31261 allocs/op
BenchmarkDecoder_Generic_JsonIter-12                        5000           4260781 ns/op         148.22 MB/s     2427840 B/op      45043 allocs/op
BenchmarkDecoder_Generic_GoJson-12                          5000           4088162 ns/op         154.47 MB/s     2757799 B/op      39700 allocs/op
BenchmarkDecoder_Generic_Sonic-12                           5000           2561050 ns/op         246.58 MB/s     2609818 B/op      12137 allocs/op
BenchmarkDecoder_Generic_Sonic_V1-12                        5000           3041447 ns/op         207.64 MB/s     2344109 B/op      29781 allocs/op
BenchmarkDecoder_Generic_Sonic_Fast-12                      5000           2488119 ns/op         253.81 MB/s     1967848 B/op      12136 allocs/op
BenchmarkDecoder_Generic_Sonnet-12                          5000           3212028 ns/op         196.61 MB/s     2096696 B/op      30183 allocs/op
BenchmarkDecoder_Parallel_Generic_StdLib-12                 5000           1534652 ns/op         411.50 MB/s     2152041 B/op      31262 allocs/op
BenchmarkDecoder_Parallel_Generic_JsonIter-12               5000           1197112 ns/op         527.53 MB/s     2427316 B/op      45043 allocs/op
BenchmarkDecoder_Parallel_Generic_GoJson-12                 5000           1046451 ns/op         603.48 MB/s     2756561 B/op      39698 allocs/op
BenchmarkDecoder_Parallel_Generic_Sonic-12                  5000            699298 ns/op         903.07 MB/s     2614464 B/op      12136 allocs/op
BenchmarkDecoder_Parallel_Generic_Sonic_V1-12               5000            906161 ns/op         696.91 MB/s     2343733 B/op      29780 allocs/op
BenchmarkDecoder_Parallel_Generic_Sonic_Fast-12             5000            784229 ns/op         805.27 MB/s     1967387 B/op      12135 allocs/op
BenchmarkDecoder_Parallel_Generic_Sonnet-12                 5000           1057777 ns/op         597.02 MB/s     2096632 B/op      30182 allocs/op
BenchmarkDecoder_Binding_StdLib-12                          5000           5519137 ns/op         114.42 MB/s      601858 B/op       5848 allocs/op
BenchmarkDecoder_Binding_JsonIter-12                        5000           1831864 ns/op         344.74 MB/s      748192 B/op      18781 allocs/op
BenchmarkDecoder_Binding_GoJson-12                          5000           1178029 ns/op         536.08 MB/s      901550 B/op       2799 allocs/op
BenchmarkDecoder_Binding_Sonic-12                           5000           1199569 ns/op         526.45 MB/s     1105079 B/op       1683 allocs/op
BenchmarkDecoder_Binding_Sonic_V1-12                        5000           1232027 ns/op         512.58 MB/s      560149 B/op       4533 allocs/op
BenchmarkDecoder_Binding_Sonic_Fast-12                      5000           1125198 ns/op         561.25 MB/s      464553 B/op       1682 allocs/op
BenchmarkDecoder_Binding_Sonnet-12                          5000           3055920 ns/op         206.65 MB/s     2096695 B/op      30183 allocs/op
BenchmarkDecoder_Parallel_Binding_StdLib-12                 5000           1182782 ns/op         533.92 MB/s      601871 B/op       5848 allocs/op
BenchmarkDecoder_Parallel_Binding_JsonIter-12               5000            476816 ns/op        1324.44 MB/s      748068 B/op      18781 allocs/op
BenchmarkDecoder_Parallel_Binding_GoJson-12                 5000            271643 ns/op        2324.79 MB/s      901943 B/op       2795 allocs/op
BenchmarkDecoder_Parallel_Binding_Sonic-12                  5000            281690 ns/op        2241.87 MB/s     1108846 B/op       1683 allocs/op
BenchmarkDecoder_Parallel_Binding_Sonic_V1-12               5000            296465 ns/op        2130.15 MB/s      560588 B/op       4533 allocs/op
BenchmarkDecoder_Parallel_Binding_Sonic_Fast-12             5000            264516 ns/op        2387.44 MB/s      465060 B/op       1682 allocs/op
BenchmarkDecoder_Parallel_Binding_Sonnet-12                 5000            359303 ns/op        1757.61 MB/s      396600 B/op       4348 allocs/op
BenchmarkEncoder_Generic_StdLib-12                          5000           5943144 ns/op         106.26 MB/s     1920483 B/op      31055 allocs/op
BenchmarkEncoder_Generic_JsonIter-12                        5000           2492510 ns/op         253.36 MB/s      637167 B/op       3793 allocs/op
BenchmarkEncoder_Generic_GoJson-12                          5000           3550643 ns/op         177.86 MB/s     1100084 B/op        544 allocs/op
BenchmarkEncoder_Generic_Sonic-12                           5000           1134123 ns/op         556.83 MB/s      470115 B/op          5 allocs/op
BenchmarkEncoder_Generic_Sonic_V1-12                        5000           1649071 ns/op         382.95 MB/s      709381 B/op          5 allocs/op
BenchmarkEncoder_Generic_Sonic_Fast-12                      5000           1147522 ns/op         550.33 MB/s      470639 B/op          5 allocs/op
BenchmarkEncoder_Parallel_Generic_Sonic-12                  5000            240012 ns/op        2631.17 MB/s      472896 B/op          4 allocs/op
BenchmarkEncoder_Parallel_Generic_JsonIter-12               5000            485174 ns/op        1301.62 MB/s      644095 B/op       3793 allocs/op
BenchmarkEncoder_Parallel_Generic_GoJson-12                 5000            749546 ns/op         842.53 MB/s     1060605 B/op        550 allocs/op
BenchmarkEncoder_Parallel_Generic_StdLib-12                 5000           1702450 ns/op         370.94 MB/s     1922308 B/op      31055 allocs/op
BenchmarkEncoder_Parallel_Generic_Sonic_V1-12               5000            582658 ns/op        1083.85 MB/s      712830 B/op          5 allocs/op
BenchmarkEncoder_Parallel_Generic_Sonic_Fast-12             5000            348096 ns/op        1814.19 MB/s      471770 B/op          4 allocs/op
BenchmarkEncoder_Binding_StdLib-12                          5000           1115137 ns/op         566.31 MB/s      324698 B/op       1423 allocs/op
BenchmarkEncoder_Binding_JsonIter-12                        5000           1087129 ns/op         580.90 MB/s      275567 B/op        314 allocs/op
BenchmarkEncoder_Binding_GoJson-12                          5000            537630 ns/op        1174.63 MB/s      265792 B/op         15 allocs/op
BenchmarkEncoder_Binding_Sonic-12                           5000            212572 ns/op        2970.83 MB/s      264782 B/op          4 allocs/op
BenchmarkEncoder_Binding_Sonic_V1-12                        5000            255639 ns/op        2470.34 MB/s      391700 B/op          4 allocs/op
BenchmarkEncoder_Binding_Sonic_Fast-12                      5000            216416 ns/op        2918.05 MB/s      265716 B/op          4 allocs/op
BenchmarkEncoder_Parallel_Binding_StdLib-12                 5000            186556 ns/op        3385.12 MB/s      325774 B/op       1423 allocs/op
BenchmarkEncoder_Parallel_Binding_JsonIter-12               5000            174684 ns/op        3615.17 MB/s      278099 B/op        314 allocs/op
BenchmarkEncoder_Parallel_Binding_GoJson-12                 5000            115085 ns/op        5487.37 MB/s      268657 B/op         15 allocs/op
BenchmarkEncoder_Parallel_Binding_Sonic-12                  5000             43834 ns/op        14406.79 MB/s     267855 B/op          4 allocs/op
BenchmarkEncoder_Parallel_Binding_Sonic_V1-12               5000             63657 ns/op        9920.59 MB/s      394323 B/op          4 allocs/op
BenchmarkEncoder_Parallel_Binding_Sonic_Fast-12             5000             45873 ns/op        13766.56 MB/s     267750 B/op          4 allocs/op
```
