## Develop an API that moves a rover around a grid

#### Rules

- You are given the initial starting point (0, 0, N) of a rover
- 0,0 are X,Y coordinates on a grid of (10, 10)
- N is the direction it is facing (N, S, E, W)
- L and R allow the rover to rotate left and right around the grid
- M allows the rover to move one point in the current direction
- The rover receives a char array of commands e.g. RMMLM and returns the finishing point after the moves e.g. 2:1:N
- The rover wraps around if it  reaches the end of the grid
- The grid may have obstacles. if a given sequence of commands encounters an obstacle, the rover moves up to the last position and reports the obstacle e.g. O:2:2:N
