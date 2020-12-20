<template>
    <div class="position-absolute bottom-0 end-x">
    <div class="card wid mb-3 hei">
        <div class="card-body overflow-y courier">
            <p class="text-center title">{{ artTitle }}</p>
            <p v-for="item in artContents" :class="defineClass(item)" :key="randomKey(item)">{{ item }}</p>
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
            // Only numbers
            if ( /[0-9]{1,}$/.test(item) ) {
                spanClass = "text-center title"
            }

            // Only Capitalized letters
            if ( /[A-Z]{1,}$/.test(item) ) {
                spanClass = "weight-800"
            }
            
            // Only Capitalized letters with "." at the final
            if ( /[A-Z]{1,}[.]$/.test(item) ) {
                spanClass = "weight-600"
            }

            // Match string where starts withs "[_" and ends with "_]"
            if ( /[[_]{2}.{1,}[\]_]{2}$/.test(item) ) {
                spanClass = "weight-600"
            }
            return spanClass 
        },
        randomKey(key) {
            return `${Math.random()}-${key}` 
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

    .weight-600 {
        font-weight: 600 !important;
    }
    
    .weight-700 {
        font-weight: 700 !important;
    }
    
    .weight-800 {
        font-weight: 800 !important;
    }
    
    .weight-900 {
        font-weight: 900 !important;
    }
    
    .overflow-y {
        overflow-y: scroll !important;
    }
    .end-x {
        left: 58vw !important;
    }
</style>
