set terminal wxt persist size 640,384
set object 1 rectangle from screen 0,0 to screen 1,1 fillcolor rgb 'white' behind
unset xtics
unset ytics
unset border
set label " 75" l rotate by 0 at 0.600,0.500 front
set label " One" l rotate by 0 at 0.800,0.000 front
set label " Two" l rotate by 0 at 0.900,1.000 front
set label " 69" l rotate by 0 at 0.500,2.500 front
set label " Three" l rotate by 0 at 1.000,2.000 front
set label " Four" l rotate by 0 at 0.800,3.000 front
set label " Five" l rotate by 0 at 0.700,4.000 front
set label "0.1" c rotate by 0 at 0.950,4.600 front
set title "newick_1"
plot "-" t "" w l lc "black"
0.000 2.750
0.000 1.500

0.300 1.500
0.000 1.500

0.300 1.500
0.300 0.500

0.600 0.500
0.300 0.500

0.600 0.500
0.600 0.000

0.800 0.000
0.600 0.000

0.600 0.500
0.600 1.000

0.900 1.000
0.600 1.000

0.300 1.500
0.300 2.500

0.500 2.500
0.300 2.500

0.500 2.500
0.500 2.000

1.000 2.000
0.500 2.000

0.500 2.500
0.500 3.000

0.800 3.000
0.500 3.000

0.000 2.750
0.000 4.000

0.700 4.000
0.000 4.000

1.000 4.400
0.900 4.400

0.950 4.600
0.950 4.600

1.200 0
