<template>
    <div class="position-absolute bottom-0 end-x">
    <div class="card wid mb-2">
        <div class="card-body">
            {{ artTitle }} 
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
}
</script>

<style scoped>
    .bottom-x {
        bottom: 0% !important;
    }
    .wid {
        width: 40vw !important;
    }
    .end-x {
        left: 58vw !important;
    }
</style>
