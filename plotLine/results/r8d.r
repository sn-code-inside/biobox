library("ggplot2")
data <- read.table(file="stdin")
plot <- ggplot(data, aes(V1, V2, group=V3))

plot <- plot + xlab(NULL)


plot <- plot + ylab(NULL)



plot <- plot + geom_path(aes(color=factor(V3)))

plot <- plot + labs(color="")

quartz(width=2.3622047244094486)
plot(plot)
while(names(dev.cur()) != 'null device')
    Sys.sleep(0.1)