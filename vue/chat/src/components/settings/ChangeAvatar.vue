<template>
    <form @submit.prevent>
        <div class="form-group">
            <label for="exampleFormControlFile1">Chose new profile photo</label>
            <input @change = "file = $event.target.files[0]" type="file" accept="multipart/form-data" class="form-control-file" id="exampleFormControlFile1">
        </div>
        <button  @click = "ChangeAvatar" type="button" class="btn btn-primary b">Update</button>
    </form>
</template>

<script>
import axios from 'axios'
export default {
    props: [
        'jwt',
    ],

    data() {
        return {
            file: null,
        }
    },

    methods: {
        async ChangeAvatar() {
            const config = { headers: { 'Content-Type': 'multipart/form-data', Authorization: 'Bearer' + ' ' + this.jwt } };
            let fd = new FormData();
            fd.append('file', this.file)
            const response = await axios.post("http://localhost:8000/a/changeavatar", fd, config)
            .catch(function (error) {
                if (error.response) {
                    alert("error")
                }
            })

            /*const response = await axios({
                method: 'post',
                url: 'http://localhost:8000/a/changeavatar',
                headers: {
                    'Content-Type': 'multipart/form-data',
                    Authorization: 'Bearer' + ' ' + this.jwt,
                },
                data: {
                    file: this.file,
                }
            })
            .catch(function (error) {
                if (error.response) {
                    alert("error")
                }
            })*/

            console.log(response)
            if (response != undefined) {
                alert("success")
            } 
        },
        handleFileUpload() {
            this.file = this.$refs.file.files[0];
        }
    }
}
</script>

<style scoped>
.b {
    margin-top: 1%;
    margin-bottom: 1%;
}
</style>