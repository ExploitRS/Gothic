<script lang="ts">
    import Nav from '../lib/components/Nav/Nav.svelte';
    import { base } from '$app/paths';
</script>

<Nav />
<nav>
    <a href="{base}/">Home</a>
    <a href="{base}/about">About</a>
</nav>
<slot />
