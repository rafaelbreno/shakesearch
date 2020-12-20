<template>
    <nav class="navbar navbar-expand-lg navbar-light bg-transparent fixed-top">
        <div class="container-fluid">
            <div class=" shake-brand px-2 mt-1">
                Shakesearch
            </div>
        </div>
    </nav>    
</template>

<script>
export default {
    name: 'Navbar'    
}
</script>

<style scoped>
    @import url('https://fonts.googleapis.com/css2?family=Amethysta&display=swap');

    .shake-brand {
        font-family: 'Amethysta', serif;
        font-weight: 400;
        font-size: 1.1rem;
        color: rgba(0, 0, 0, 0.6) !important;
    }

</style>
