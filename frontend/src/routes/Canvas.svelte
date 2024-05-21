<script lang="ts">
    import { onMount } from 'svelte';
    import { dia, shapes, util } from '@joint/core';
    import { dndzone } from 'svelte-dnd-action';

    /**
	 * @type {any}
	 */
    let graph;
    let paper;

    onMount(() => {
        graph = new dia.Graph();
        paper = new dia.Paper({
            el: document.getElementById('canvas'),
            model: graph,
            width: 800,
            height: 600,
            gridSize: 10,
            drawGrid: true
        });

        // Define the custom heptagon shape
        shapes.custom = {};
        shapes.custom.Heptagon = shapes.standard.Polygon.extend({
            markup: '<g class="rotatable"><g class="scalable"><polygon class="body"/></g><text class="label"/></g>',
            defaults: util.deepSupplement({
                type: 'custom.Heptagon',
                attrs: {
                    'polygon.body': {
                        refPoints: '0,30 24,2 76,2 100,30 81,68 19,68 0,30',
                        fill: 'lightblue',
                        stroke: 'black',
                        strokeWidth: 2
                    },
                    'text.label': {
                        text: 'Heptagon',
                        refX: '50%',
                        refY: '50%',
                        yAlignment: 'middle',
                        xAlignment: 'middle',
                        fill: 'black',
                        fontSize: 12,
                        fontWeight: 'bold'
                    }
                }
            }, shapes.standard.Polygon.prototype.defaults)
        });
    });

    const addElement = (icon) => {
        const element = new shapes.custom.Heptagon();
        element.position(100, 100);
        element.resize(100, 100);
        element.attr('text.label/text', icon.name);
        element.addTo(graph);
    };

    let icons = [
        { id: 1, name: 'Private Endpoint', src: 'path/to/icon.png' },
        { id: 2, name: 'Public Endpoint', src: 'path/to/icon.png' }
    ];

    const handleDrop = (event) => {
        const { items } = event.detail;
        // Handle drop event here
    };
</script>

<style>
    #canvas {
        border: 1px solid #ccc;
    }

    .palette {
        display: flex;
        flex-direction: column;
        width: 200px;
        border: 1px solid #ccc;
        padding: 10px;
    }
</style>

<div class="container">
    <div class="palette" use:dndzone={{ items: icons, type: 'icon', onDrop: handleDrop }}>
        {#each icons as icon}
            <div class="icon" data-id={icon.id}>
                <img src={icon.src} alt={icon.name} />
                <span>{icon.name}</span>
            </div>
        {/each}
    </div>
    <div id="canvas"></div>
</div>
