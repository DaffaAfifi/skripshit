const cpuRamAvg = () => {
  const data = `1.50 2.80
1.50 2.80
1.50 2.80
1.50 2.80
1.50 2.80
1.50 2.80
1.50 2.80
1.50 2.80
1.50 2.80
1.60 2.80
1.70 2.80
1.70 2.80
1.80 2.80
1.90 2.80
2.00 2.80
2.10 2.80
2.20 2.80
2.20 2.80
2.20 2.80`
    .split("\n")
    .map((line) => {
      const [cpu, ram] = line.split(" ").map(Number);
      return { cpu, ram };
    });

  const avgCpu = data.reduce((sum, item) => sum + item.cpu, 0) / data.length;
  const avgRam = data.reduce((sum, item) => sum + item.ram, 0) / data.length;

  return { avgCpu, avgRam };
};

const result = cpuRamAvg();
console.log(result.avgCpu, result.avgRam);
