<template>
    <nav class="navbar navbar-expand-lg navbar-light bg-transparent fixed-top">
        <div class="container-fluid">
            <span href="#" 
               class="navbar-brand shake-brand px-2 mt-1" 
               disabled>
                Shakesearch
            </span>
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
        opacity: 40%;
    }

</style>
