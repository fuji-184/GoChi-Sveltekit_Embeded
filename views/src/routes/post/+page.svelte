<script lang="ts">
        import { 
                
                Table,
                tableMapperValues,
                Paginator
                
        } from '@skeletonlabs/skeleton'
        


const sourceData = [{ position: 1, name: 'Hydrogen', weight: 1.0079, symbol: 'H' },
	{ position: 2, name: 'Helium', weight: 4.0026, symbol: 'He' },
	{ position: 3, name: 'Lithium', weight: 6.941, symbol: 'Li' },
	{ position: 4, name: 'Beryllium', weight: 9.0122, symbol: 'Be' },
	{ position: 5, name: 'Boron', weight: 10.811, symbol: 'B' },
		{ position: 1, name: 'Hydrogen', weight: 1.0079, symbol: 'H' },
	{ position: 2, name: 'Helium', weight: 4.0026, symbol: 'He' },
	{ position: 3, name: 'Lithium', weight: 6.941, symbol: 'Li' },
	{ position: 4, name: 'Beryllium', weight: 9.0122, symbol: 'Be' },
	{ position: 5, name: 'Boron', weight: 10.811, symbol: 'B' }];

let paginationSettings = {
	page: 0,
	limit: 5,
	size: sourceData.length,
	amounts: [1,2,5,10,20],
}
				
$: paginatedSource = sourceData.slice(
	paginationSettings.page * paginationSettings.limit,
	paginationSettings.page * paginationSettings.limit + paginationSettings.limit
);				
				
const tableSimple = {
	// A list of heading labels.
	head: ['Name', 'Symbol', 'Weight'],
	// The data visibly shown in your table body UI.
	body: tableMapperValues(sourceData, ['name', 'symbol', 'weight']),
	

	foot: ['Total', '', '<code class="code">5</code>']
};

function mySelectionHandler () {
       console.log(100) 
}			



$: fetchData = {}

async function onPageChange(e: CustomEvent): void {
  try {
    fetchData = await (await fetch(`http://localhost:8082/post/${e.detail}`)).json()
    console.log(fetchData)
  } catch (error) {
    console.error("Gagal mengambil data:", error)
  }
}

function onAmountChange(e: CustomEvent): void {
	console.log('event:amount', e.detail);
}
			

					
</script>

<Table source={tableSimple} interactive={true} on:selected={mySelectionHandler} />

<Paginator
	bind:settings={paginationSettings}
	showFirstLastButtons="{false}"
	showPreviousNextButtons="{true}" on:page={onPageChange} on:amount={onAmountChange}
/>

{#if fetchData.length > 0}
    <ul>
      {#each fetchData as item (item.no)}
        <li>No: {item.no}, Data: {item.data}</li>
      {/each}
    </ul>
  {:else}
    <p>Data belum diambil.</p>
  {/if}
