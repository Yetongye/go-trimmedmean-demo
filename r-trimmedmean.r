set.seed(123)
float_data <- rnorm(100, mean=50, sd=10)
int_data <- sample(0:99, 100, replace=TRUE)

cat("R Trimmed Mean of floats: ", mean(float_data, trim=0.05), "\n")
cat("R Trimmed Mean of ints:   ", mean(int_data, trim=0.05), "\n")


set.seed(123)
data <- rnorm(100, mean=50, sd=10)
B <- 1000
means <- numeric(B)
tmeans <- numeric(B)

for (b in 1:B) {
  sample <- sample(data, 100, replace=TRUE)
  means[b] <- mean(sample)
  tmeans[b] <- mean(sample, trim=0.1)
}

cat("Bootstrap Mean:         ", mean(means), " (SE =", sd(means), ")\n")
cat("Bootstrap Trimmed Mean: ", mean(tmeans), " (SE =", sd(tmeans), ")\n")
