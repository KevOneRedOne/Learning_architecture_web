<script setup onload>

  const { pending, data: articles} = useFetch('http://localhost:8080/api/v1/articles/', {
    headers: {
      'Content-Type': 'application/json',
      'Accept': 'application/json'
    }
  });

  
</script>



<template>
 <div class="container">
      <div class="grid">    
          <div v-for="article in articles">
            <div class="item">
              <h2>{{ article.title }}</h2>
              <p>{{ article.description }}</p>
              <p>Date : {{ article.date }}</p>
            </div>
          </div>
      </div>
  </div>
</template>


