<template>
  <apexchart type="radialBar" :options="options" :series="series"></apexchart>
</template>

<script>
export default {
  data() {
    return {
      series: [0],
      options: {
        labels: ['CPU Usage']
      }
    };
  },
  mounted: function() {
    // eslint-disable-next-line no-undef
    wails.Events.On("usage", cpu_usage => {
      if (cpu_usage) {
        this.series = [ cpu_usage.avg ];
      }
    });
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1 {
  margin-top: 2em;
  position: relative;
  min-height: 5rem;
  width: 100%;
}
a:hover {
  font-size: 1.7em;
  border-color: blue;
  background-color: blue;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
  transition: 500ms;
}
a {
  font-size: 1.7em;
  border-color: white;
  background-color: #121212;
  color: white;
  border: 3px solid white;
  border-radius: 10px;
  padding: 9px;
  cursor: pointer;
}
</style>
