<template>
    <div class = "message-container">
        <slot name = "image"></slot>
        <p><slot name = "message"></slot></p>
        <span class="time"><slot name = "time"></slot></span>
    </div>
</template>

<style scoped>
.message-container {
    border: 2px solid #ccc;
    background-color: #ddd;
    text-align: left;
    border-radius: 5px;
    padding: 30px;
    margin: 10px 0;
    bottom:0;
}

.time {
    float: left;
    color: #999;
}
</style>