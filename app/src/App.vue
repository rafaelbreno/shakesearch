<template>
    <div class="container-fluid overflow-hidden">
        <Navbar/>
        <div class="position-relative">
            <SearchBar/>
            <Shakespeare/>
        </div>
    </div>
</template>

<script>
import Navbar from './components/_partials/Navbar.vue'
import SearchBar from './components/SearchBar.vue'
import Shakespeare from './components/Shakespeare.vue'

export default {
    name: 'App',
    components: {
        Navbar,
        SearchBar,
        Shakespeare
    }
}
</script>

<style>

</style>
