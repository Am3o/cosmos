<template>
  <div>
    <router-view />
  </div>
</template>

<style>
.sp-container {
  margin: 0 auto;
  max-width: 800px;
  padding: 1rem;
}
</style>

<script>
export default {
  created() {
    this.$store.dispatch("cosmos/init");
    this.$store.dispatch("cosmos/entityFetch", {type: "poll", module: "cosmos"});
    this.$store.dispatch("cosmos/entityFetch", {type: "vote", module: "cosmos"});
  },
};
</script>
