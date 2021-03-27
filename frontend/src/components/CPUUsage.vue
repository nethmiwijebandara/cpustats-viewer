<template>
<div>

 <h1>{{avg}}</h1>
  <apexchart type="radialBar" :options="options" :series="series"></apexchart>
  </div>
 </template>

<script>
export default {
  data() {
    return {
      series: [],
      avg:'',
      options: {
        labels: ['CPU Usage']
      }
    };
  },
  mounted: function() {
    // eslint-disable-next-line no-undef
    wails.Events.On("usage", cpu_usage => {
      console.log(cpu_usage, "usage")
      console.log(this.series, "series")
      if (cpu_usage) {
        this.series =  cpu_usage.avg ;
        this.avg = cpu_usage.avg;
      }
    });
  }
};
</script>