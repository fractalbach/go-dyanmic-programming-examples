# Staircase Problem

Open the .go file and read the comments to see more.

Here is example output from calculating 
the number of ways to climb a staircase with 100 steps (`n = 100`)
With memoization using a hashmap, it's pretty fast:
~~~
RunTime: 89.737µs
~~~

## Solutions

~~~
f(0) = 0
f(1) = 1
f(2) = 2
f(3) = 4
f(4) = 7
f(5) = 13
f(6) = 24
f(7) = 44
f(8) = 81
f(9) = 149
f(10) = 274
f(11) = 504
f(12) = 927
f(13) = 1705
f(14) = 3136
f(15) = 5768
f(16) = 10609
f(17) = 19513
f(18) = 35890
f(19) = 66012
f(20) = 121415
f(21) = 223317
f(22) = 410744
f(23) = 755476
f(24) = 1389537
f(25) = 2555757
f(26) = 4700770
f(27) = 8646064
f(28) = 15902591
f(29) = 29249425
f(30) = 53798080
f(31) = 98950096
f(32) = 181997601
f(33) = 334745777
f(34) = 615693474
f(35) = 1132436852
f(36) = 2082876103
f(37) = 3831006429
f(38) = 7046319384
f(39) = 12960201916
f(40) = 23837527729
f(41) = 43844049029
f(42) = 80641778674
f(43) = 148323355432
f(44) = 272809183135
f(45) = 501774317241
f(46) = 922906855808
f(47) = 1697490356184
f(48) = 3122171529233
f(49) = 5742568741225
f(50) = 10562230626642
f(51) = 19426970897100
f(52) = 35731770264967
f(53) = 65720971788709
f(54) = 120879712950776
f(55) = 222332455004452
f(56) = 408933139743937
f(57) = 752145307699165
f(58) = 1383410902447554
f(59) = 2544489349890656
f(60) = 4680045560037375
f(61) = 8607945812375585
f(62) = 15832480722303616
f(63) = 29120472094716576
f(64) = 53560898629395777
f(65) = 98513851446415969
f(66) = 181195222170528322
f(67) = 333269972246340068
f(68) = 612979045863284359
f(69) = 1127444240280152749
f(70) = 2073693258389777176
f(71) = 3814116544533214284
f(72) = 7015254043203144209
f(73) = 12903063846126135669
f(74) = 23732434433862494162
f(75) = 43650752323191774040
f(76) = 80286250603180403871
f(77) = 147669437360234672073
f(78) = 271606440286606849984
f(79) = 499562128250021925928
f(80) = 918838005896863447985
f(81) = 1690006574433492223897
f(82) = 3108406708580377597810
f(83) = 5717251288910733269692
f(84) = 10515664571924603091399
f(85) = 19341322569415713958901
f(86) = 35574238430251050319992
f(87) = 65431225571591367370292
f(88) = 120346786571258131649185
f(89) = 221352250573100549339469
f(90) = 407130262715950048358946
f(91) = 748829299860308729347600
f(92) = 1377311813149359327046015
f(93) = 2533271375725618104752561
f(94) = 4659412488735286161146176
f(95) = 8569995677610263592944752
f(96) = 15762679542071167858843489
f(97) = 28992087708416717612934417
f(98) = 53324762928098149064722658
f(99) = 98079530178586034536500564
f(100) = 180396380815100901214157639
~~~
