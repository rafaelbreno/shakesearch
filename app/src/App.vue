<template>
    <div class="container-fluid overflow-hidden">
        <Navbar/>
        <div class="position-relative">
            <SearchBar/>
        </div>
    </div>
</template>

<script>
import Navbar from './components/_partials/Navbar.vue'
import SearchBar from './components/SearchBar.vue'

export default {
    name: 'App',
    components: {
        Navbar,
        SearchBar,
    }
}
</script>

<style>

</style>
