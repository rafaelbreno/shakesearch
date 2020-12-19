<template>
    <div class="position-absolute bottom-0 end-x">
    <div class="card wid mb-3 hei">
        <div class="card-body overflow-y courier">
            <p class="text-center title">{{ artTitle }}</p>
            <p v-for="item in artContents" :class="defineClass(item)" :key="item">{{ item }}</p>
        </div>
    </div>
    </div>
</template>

<script>
export default {
    name: 'Book',
    props: [
        'artTitle',
        'artContents'
    ],
    methods: {
        defineClass(item) {
            let spanClass = ''
            if ( /[0-9]{1}$/.test(item) ) {
                spanClass = "text-center title"
            }

            if ( /[A-B]{1,}$/.test(item) ) {
                spanClass = "title"
            }
            return spanClass 
        }
    }
}
</script>

<style scoped>
    .bottom-x {
        bottom: 0% !important;
    }
    .hei {
        height: 96vh !important;
    }
    .wid {
        width: 40vw !important;
    }

    .courier {
        font-family: 'Courier New', monospace;
    }

    .title {
        font-weight: 500 !important;
    }

    .overflow-y {
        overflow-y: scroll !important;
    }
    .end-x {
        left: 58vw !important;
    }
</style>
