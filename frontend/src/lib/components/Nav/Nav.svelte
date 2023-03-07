<nav>
    <div class="logo">
        Home
    </div>
    <div class="items">
        <div>item</div>
        <div>item</div>
    </div>
</nav>

<style lang="scss">
    nav {
		display: grid;
		grid-template-columns: 1fr 1fr 1fr;
		padding-inline: 0.75em;
		position: fixed;
		top: 0;
		left: 0;
		right: 0;
		width: 100%;
		min-width: 100vw;
		grid-template-areas: "l m r";
		height: var(--top-bar-height);
		align-content: center;
		align-items: center;
		touch-action: none;
		z-index: 150;
		background-color: var(--top-bg);
		// isolation: isolate;
		justify-content: space-between;
		border-bottom: 0.0625rem hsl(0deg 0% 12%) solid;
		@media screen and (min-width: 720px) {
			z-index: 155;
		}
	}
</style>