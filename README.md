# _*Linear Stats Project*_

### _*Author*_

[AnxielRay](https://www.linkedin.com/mynetwork/grow/)

## _*Project Overview*_

- The Linear Stats Project is a Go program designed to read data from a file and compute two key statistical metrics:

    - Linear Regression Line: Computes the best-fit line for the given data points.
    - Pearson Correlation Coefficient: Measures the strength and direction of the linear relationship between the data points.

- The program reads the data from a file where each line represents a y-value, and the x-values are implicitly the line numbers starting from 0.
Features

    - `Linear Regression Line Calculation`: Computes the equation of the line in the form y = ax + b where a is the slope and b is the y-intercept.
    - `Pearson Correlation Coefficient`: Calculates the correlation coefficient with 10 decimal places of precision.
    - `Command-Line Interface`: Accepts the file path as an argument.

### Requirements

    `Go 1.18 or later`

## _*Installation*_

### _*Clone the Repository:*_

```sh
git clone https://learn.zone01kisumu.ke/git/rogwel/linear-stats.git
cd linear-stats
```

- Download the required file to test the program, then extract it.
- Move the source code and the module file into the `stats-bin-dockerized/stat-bin`

### _*Run the file*_

- Run the proram script first using the command:

```sh
./bin/linear-stats
```

This will generate a data textual file, that our program should run against.

## _*Testing*_

- When we run our program:

```sh
go run . data.txt
```

- The program output should give an output equal to the one from the script that we executed.
- The program will output the results calculated, of the linear regression line and the Peearsons Correlation coefficient in different lines:

```sh
Linear Regression Line: y = 1.234567x + 45.678901
Pearson Correlation Coefficient: 0.9876543210
```

### _*Functions*_

    - `main`: Reads the file path from the command line, processes the data, and prints the results.
    - `InterceptAndSlope(data string) (slope, intercept float64):` Calculates the slope and intercept of the linear regression line.
    - `correlation(data string) float64:` Computes the Pearson Correlation Coefficient.
    - `sum(data []float64) float64:` Calculates the sum of a slice of float64 values.
    - `sumSquares(Xs []float64) float64:` Calculates the sum of the squares of a slice of float64 values.
    - `sumProduct(Xs, Ys []float64) float64:` Calculates the sum of the product of two slices of float64 values.
    - `mean(data []float64) float64:` Computes the mean of a slice of float64 values.

### *License*

- This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### Contributing

- If you would like to contribute to this project, please fork the repository and submit a pull request. For major changes or improvements, please open an issue first to discuss what you would like to change.
Contact