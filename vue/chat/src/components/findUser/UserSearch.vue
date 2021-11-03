<template>
    <div class="input-group search">
        <input  v-model="login" type="search" class="form-control rounded" placeholder="Search" aria-label="Search" aria-describedby="search-addon" />
        <button @click = "findUser" type="button" class="btn btn-outline-primary">search</button>
    </div>
</template>

<script>
export default{
    data(){
        return{
            login:""
        }
    },
    methods: {
        findUser() {
            this.$emit("findUser", this.login);
        },
    }
}

</script>

<style scoped>
    .search {
        margin-top: 1%;
        margin-bottom: 1%;
        width:auto;
    }
</style>