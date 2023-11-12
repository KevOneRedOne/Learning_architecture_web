<script setup>
    const config = useRuntimeConfig();
    const search = ref('');
    const results = ref([]);

    watchEffect(() => {
        console.log(search.value);
        if (search.value.length > 2) {
            searchArticles(search.value).then((data) => {
                results.value = data;
                console.log(results.value);
            });
        } else {
            results.value = [];
        }
    },[search]);


    const searchArticles =  async (searchValue) => {
        console.log(searchValue);
        const { data: articles } =  await useFetch(`http://127.0.0.1:8181/api/search?title=${searchValue}`, {
            onResponse({ request, response }) {
                console.log("[fetch response]", request, response.status, response.body);
            },
            onRequestError({ request, error }) {
                console.error("[fetch request error]", request, error);
            },
        });
        return articles.value;
    };


</script>

<template>
    <div class="wrapper">
        <input 
            type="text" 
            placeholder="Rechercher un article" 
            v-model="search"
        />
        <div class="results">
            <p v-for="result in results" :key="result.id">{{ result.title}}</p><br>
        </div>
    </div>

</template>

<style scoped>

.wrapper {
    display: flex;
    border: 1px solid #12b488;
    padding: 7px;
    border-radius: 4px;
    position: relative;
    background-color: #fff;
}

.wrapper svg{
    width: 20px;
    margin-right: 10px;
}

.wrapper input {
    border: none;
}
.wrapper input:focus {
    outline: none;
}
.results {
    position: absolute;
    max-height: 150px;
    left: 0px;
    top: 35px;
    background-color: #fff;
    padding: 10px;
    width: 100%;
    border: solid 1px #333;
    border-radius: 5px;
    display: none;
}

.wrapper input:focus + .results {
    display: block;
}

.results p {
    font-size: 12px;
    color: #333;
}

.results p:hover {
    color: #12b488 !important;
}


</style>