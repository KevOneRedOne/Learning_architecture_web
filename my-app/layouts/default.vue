<template>
  <div>
    <header class="main-header">
        <NuxtLink to="/"><h1>My Wonderful App</h1></NuxtLink>
        <nav>
          <ul class="navbar">
            <li><NuxtLink to="/create">Créer un article</NuxtLink></li>
            <li><NuxtLink to="/articles" class="btn">Articles</NuxtLink></li>
            <li><SearchInput /></li> 
          </ul>
        </nav>
      
        
    </header>

    <div class="container">
      <slot />
    </div>
  </div>
</template>
  
<style scoped>
  .main-header {
    background-color: #333;
    color: #fff;
    padding: 1rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .navbar {
    background-color: #333;
    color: #fff;
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem;
    position: relative;
  }

  h1 {
    text-transform: uppercase;
    letter-spacing: 1px;
    font-size: 15px;
    padding: 1rem;
  }

  ul {
    display: flex;
    list-style: none;
    margin: 0;
    padding: 0;
  }

  li {
    margin-left: 1rem;
  }

  a {
    color: #fff;
    text-decoration: none;
  }

  .router-link-exact-active {
    font-weight: bold;
    color: #12b488;
  }

  .btn {
    background-color: #12b488;
    color: #fff;
    padding: 0.5rem 1rem;
    border-radius: 5px;
    border: 1px solid #12b488;
  }

  .btn:hover {
    background-color: #fff;
    color: #12b488;
    padding: 0.5rem 1rem;
    border: 1px solid #12b488;
    border-radius: 5px;
  }

  .btn.router-link-exact-active {
    background-color: #fff;
    color: #12b488;
    padding: 0.5rem 1rem;
    border: 1px solid #12b488;
    border-radius: 5px;
  }

</style>
  