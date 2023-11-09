<script setup onload>
  const config = useRuntimeConfig();
  if (!config.public.api) {
    console.error('API_PATH is not defined. Please check your environment configuration.');
    console.log(config.public.api)
  } else {
    const { pending, data: articles } = useFetch(`${config.public.api}/articles/`, {
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json'
      }
    });

    console.log(articles);
  }
</script>

<template>
  <div class="container">
    <div class="grid">
      <div v-if="!pending" v-for="article in articles" :key="article.id">
        <div class="item">
          <h2>{{ article.title }}</h2>
          <p>{{ article.description }}</p>
          <p>Date : {{ article.date }}</p>
        </div>
      </div>
      <div v-if="pending">Chargement en cours...</div>
    </div>
  </div>
</template>
