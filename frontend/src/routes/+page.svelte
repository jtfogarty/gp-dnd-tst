<script lang="ts">
    import IconPalette from './IconPalette.svelte';
    import Canvas from './Canvas.svelte';

    let icons = [
        { id: 1, name: 'Private Endpoint', src: 'path/to/icon.png' },
        { id: 2, name: 'Public Endpoint', src: 'path/to/icon.png' }
    ];
</script>

<style>
    .container {
        display: flex;
    }
</style>

<div class="container">
    <IconPalette {icons} />
    <Canvas />
</div>
