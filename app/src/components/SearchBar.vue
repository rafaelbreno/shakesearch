<template>
    <div class="position-relative vh-100">
        <div class="position-absolute top-45 start-49 translate-middle">
            <div class="row" :class="{'search-input-move': !show_shake, 'search-input-default': show_shake}">
                <div class="col-1">
                    <img src="../assets/svg/magnifying-glass.svg" 
                         class="py-1 px-2"
                         alt="sketch"/>
                </div>
                <div class="col-10">
                    <input type="text"
                           name="keywords"
                           id="keywords"
                           class="shake-search-input"
                           placeholder="What art thee looking f'r?"
                           v-model="form.keywords"/>
                </div>
                <div class="col-1">
                    <transition name="arrow-button-fade">
                        <button 
                            id="search-button" 
                            class="shake-button pb-1 me-3"
                            v-on:click="searchArt"
                            v-if="form.keywords != ''">
                            <i class="fas fa-arrow-right"></i>
                        </button>
                    </transition>
                </div>
            </div>
        </div>
        <div class="position-absolute top-50 start-49 translate-middle">
            <transition name="sketch-fade">
                <img src="../assets/svg/sketch.svg"
                     alt="sketch"
                     v-if="show_shake"/>
            </transition>
        </div>
        <transition name="shake-fade">
            <Shakespeare v-if="show_shake"/>
        </transition>
        <transition name="book-fade">
            <Book v-bind:artContents="content.body" 
                  v-bind:artTitle="content.title" 
                  v-if="!show_shake" />
        </transition>
    </div>
</template>

<script>
import Shakespeare from './Shakespeare.vue'
import Book from './_partials/Book.vue'

export default {
    name: 'SearchBar',
    components: {
        Shakespeare,
        Book
    },
    data: () => ({
        show_shake: true,
        form: {
            keywords: "",
        },
        content: {
            title: "",
            body: [],
        },
    }),
    methods: {
        searchArt() {
            let data = {
                query: this.form.keywords
            }
            this.$postData("", "POST", data)
            .then(r => {
                this.show_shake = !this.show_shake
                this.content.title = r.contents[0].Title
                this.content.body = r.contents[0].Body
                console.log(r)
            }).catch(e => {
                this.show_shake = !this.show_shake
                console.log(e)
            })
        },
    },
}
</script>

<style scoped>
    @import url('https://fonts.googleapis.com/css2?family=Courier+Prime:ital,wght@0,400;0,700;1,400;1,700&display=swap');
    /*font-family: 'Courier Prime', monospace;*/

    @import url('https://fonts.googleapis.com/css2?family=Playfair+Display:ital,wght@0,400;0,500;0,600;0,700;0,800;0,900;1,400;1,500;1,600;1,700;1,800;1,900&display=swap');
    /*font-family: 'Playfair Display', serif;*/

    .search-input-default {
        transform: translateX(0%);
        transition: transform .8s;
    }

    .search-input-move {
        transform: translateX(-40%);
        transition: transform .8s;
    }
    .arrow-button-fade-enter-active {
        transition: all .8s ease;
    }
    .arrow-button-fade-leave-active {
        transition: all .4s cubic-bezier(1.0, 0.5, 0.8, 1.0);
    }
    .arrow-button-fade-enter, .arrow-button-fade-leave-to {
    /*[> .slide-fade-leave-active below version 2.1.8 <] {*/
        transform: translateX(-10%);
        opacity: 0;
    }

    .sketch-fade-enter-active {
        transition: all .8s ease;
    }
    .sketch-fade-leave-active {
        transition: all .8s cubic-bezier(1.0, 0.5, 0.8, 1.0);
    }
    .sketch-fade-enter, .sketch-fade-leave-to
    /* .slide-fade-leave-active below version 2.1.8 */ {
        transform: translateX(-40%);
        opacity: 0;
    }

    .shake-fade-enter-active {
        transition: all .8s ease;
    }
    .shake-fade-leave-active {
        transition: all .8s cubic-bezier(1.0, 0.5, 0.8, 1.0);
    }
    .shake-fade-enter, .shake-fade-leave-to
    /* .slide-fade-leave-active below version 2.1.8 */ {
        transform: translateX(80%);
        opacity: 0;
    }

    .book-fade-enter-active {
        transition: all .8s ease;
    }
    .book-fade-leave-active {
        transition: all .8s cubic-bezier(1.0, 0.5, 0.8, 1.0);
    }
    .book-fade-enter, .book-fade-leave-to
    /* .slide-fade-leave-active below version 2.1.8 */ {
        transform: translateY(80%);
        opacity: 0;
    }

    .shake-button {
        background-color: rgba(0, 0, 0, 0);
        border: 0;
        outline: none;

        font-size: 2rem;
        color: #9B9186;

        z-index: 999;
    }

    .start-49 {
        left: 49% !important;
    }
    
    .top-45 {
        top: 45% !important;
    }
    
    .bottom-x {
        bottom: 0% !important;
    }
    .end-x {
        right: -3% !important;
    }

    input:focus {
        border: 0;
        outline: none;
    }

    .shake-search-input {
        background-color: rgba(0, 0, 0, 0);
        border: 0; 
        color: #50493E; 
        font-family: 'Courier Prime', monospace;
        font-size: 2rem;
    }

    .shake-search-input::-webkit-input-placeholder { /* Chrome/Opera/Safari */
        color: #8A8070;
        font-family: 'Playfair Display', serif;
    }
    .shake-search-input::-moz-placeholder { /* Firefox 19+ */
        color: #8A8070;
        font-family: 'Playfair Display', serif;
    }
    .shake-search-input:-ms-input-placeholder { /* IE 10+ */
        color: #8A8070;
        font-family: 'Playfair Display', serif;
    }
    .shake-search-input:-moz-placeholder { /* Firefox 18- */
        color: #8A8070;
        font-family: 'Playfair Display', serif;
    }
</style>
