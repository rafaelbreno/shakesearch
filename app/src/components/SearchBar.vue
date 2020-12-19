<template>
    <div class="position-relative vh-100">
        <div class="position-absolute top-45 start-49 translate-middle">
            <div class="row">
                <div class="col-1">
                    <img src="../assets/svg/magnifying-glass.svg" 
                         class="py-1 px-2"
                         alt="sketch"/>
                </div>
                <div class="col">
                    <input type="text"
                           name="keywords"
                           id="keywords"
                           class="shake-search-input"
                           placeholder="What art thee looking f'r?"
                           v-model="form.keywords"/>
                </div>
                <div class="col-1">
                    <button 
                        id="search-button" 
                        class="shake-button pb-1 me-3"
                        v-on:click="searchArt">
                        <i class="fas fa-arrow-right"></i>
                    </button>
                </div>
            </div>
            <img src="../assets/svg/sketch.svg" alt="sketch"/>
        </div>
    </div>
</template>

<script>
export default {
    name: 'SearchBar',
    data: () => ({
      form: {
        keywords: "",
      },
    }),
    methods: {
        searchArt() {
            let data = {
                query: this.form.keywords
            }

            this.$postData("", "POST", data)
            .then(r => {
                console.log(r)
            }).catch(e => {
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
