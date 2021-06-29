<template>
  <div class="nav">
    <div>
      <router-link to="/">Home</router-link>
      <router-link to="/Page1">Page1</router-link>
      <router-link to="/Page2">Page2</router-link>
      <router-link to="/Page3">Page3</router-link>
      <router-link to="/Page4">Page4</router-link>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Navigation',
  props: {},
};
</script>
<style lang="scss" scoped>
@import '../../responsive.scss';
.nav {
  background-color: grey;
  display: flex;
  width: 100%;
  height: 10vh;
  justify-content: space-around;
  align-content: center;
  @include tablet{
    background-color: red;
  }
  @include mobile{
    background-color: blue;
  }
}
</style>
