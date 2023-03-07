<script lang='ts'>
    // const apiUrl = (path: string) => `${import.meta.env.VITE_API_URL || ''}${path}`;
    import type { PageServerData } from "./$types";

    export let data: PageServerData;

    let { name, text } = data;

    // const getVersion = async () => {
    //     const url = apiUrl('/');
    //     const res = await fetch(url);

    //     if (!res.ok) {
    //         throw `Error while fetching data from ${url} (${res.status} ${res.statusText}).`;
    //     }

    //     const { version } = await res.json();
    //     return version;
    // };
</script>

<h1>Welcome to Exploit.RS</h1>
<div>{name}</div>
<div>{text}</div>
<!-- {#await getVersion()}
    loading...
{:then version}
    Version from Server: {version}
{:catch err}
    {err}
{/await} -->