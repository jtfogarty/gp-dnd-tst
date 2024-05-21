<script>
    import { dndzone } from 'svelte-dnd-action';

    export let icons = [];

    const handleDrop = (event) => {
        const { items } = event.detail;
        // Handle drop event here
    };
</script>

<style>
    .palette {
        display: flex;
        flex-direction: column;
        width: 200px;
        border: 1px solid #ccc;
        padding: 10px;
    }
</style>

<div class="palette" use:dndzone={{ items: icons, type: 'icon', onDrop: handleDrop }}>
    {#each icons as icon}
        <div class="icon" data-id={icon.id}>
            <img src={icon.src} alt={icon.name} />
            <span>{icon.name}</span>
        </div>
    {/each}
</div>
