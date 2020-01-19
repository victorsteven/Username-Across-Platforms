<template>
  <div id="app">
    <v-app id="inspire">
      <v-content class="mt-6">
        <v-container>
          <h3 class="text-center" style="color: indigo">Username Lookup App</h3>
          <v-card
                  class="mx-auto mt-4"
                  max-width="344"
                  outlined
          >
            <form @submit.prevent="checkUsername">
            <v-list-item>
              <v-list-item-content>
                <v-list-item-title class="align-center">
                  <div>
                    <v-col cols="12" sm="12" md="12">
                      <v-text-field
                              label="Enter username"
                              v-model="username"
                      ></v-text-field>
                    </v-col>
                  </div>
                  <div style="font-weight: bold">Choose the Application to lookup</div>
                  <v-checkbox value="twitter" v-model="checkAccounts" :label="`Twitter`"></v-checkbox>
                  <v-checkbox value="instagram" v-model="checkAccounts" :label="`Instagram`"></v-checkbox>
                  <v-checkbox value="github" v-model="checkAccounts" :label="`Github`"></v-checkbox>
                  <v-checkbox value="dev.to" v-model="checkAccounts" :label="`Dev.to`"></v-checkbox>
                  <v-checkbox value="bitbucket" v-model="checkAccounts" :label="`BitBucket`"></v-checkbox>
                </v-list-item-title>
              </v-list-item-content>
            </v-list-item>
            <div class="text-center mb-3">
              <v-btn :disabled="disabled" type="submit" class="ma-2 style-btn" tile color="indigo">
                <span v-if="loading">Checking Username</span>
                <span v-else>Check Username</span>
              </v-btn>
            </div>
            </form>
            <div>
              <div v-if="matchValues.length > 0 && !loading && !notMatched" class="pa-2 mb-2">
                <h4 class="text-center">The username matched these accounts: </h4>
                <div v-for="(result, index) in matchValues" :key="index" class="text-center">
                  <a :href="result.url" target="_blank"> {{ result.account }}</a>
                </div>
              </div>
              <div v-if="!loading && notMatched" class="text-center mb-4">
                <h4>No account match this username</h4>
              </div>
            </div>
          </v-card>

        </v-container>
      </v-content>
      <v-footer height="auto" color="indigo" dark>
        <v-layout justify-center row wrap>
          <v-flex color="indigo" dark py-3 text-xs-center text-md-center text-sm-center white--text xs12>
            <strong>Developed with ❤️  by <a target="_blank" href="https://twitter.com/stevensunflash"><span class=" style-name">@stevensunflash</span></a> &copy; {{ new Date().getFullYear() }}</strong>
          </v-flex>
        </v-layout>
      </v-footer>
    </v-app>
  </div>
</template>

<script>

  import axios from 'axios';
  import API_ROUTE from "./env";


  export default {
    data: () => ({
      drawer: null,
      checkAccounts: [],
      matchedAccount: [],
      notMatched: false,
      username: '',
      loading: false
    }),
    computed: {
      disabled() {
        return this.loading === true || this.username === "" || this.checkAccounts.length === 0
      },
      checkUrl() {
        let arr = []
        this.checkAccounts.map(val => {
          if (val === "twitter") {
            arr.push(`https://twitter.com/${this.username}`)
          } else if (val === "instagram") {
            arr.push(`https://instagram.com/${this.username}`)
          } else if (val === "github") {
            arr.push(`https://github.com/${this.username}`)
          } else if (val === "dev.to") {
            arr.push(`https://dev.to/${this.username}`)
          } else if (val === "bitbucket") {
            arr.push(`https://bitbucket.com/${this.username}`)
          }
        });
        return arr
      },
      matchValues() {
        let finalMatch = []
        this.matchedAccount.map(oneMatch => {
          if (oneMatch.includes("twitter")){
            finalMatch.push({ "url": oneMatch, "account": "Twitter"})
          } else if (oneMatch.includes("instagram")){
            finalMatch.push({ "url": oneMatch, "account": "Instagram"})
          } else if (oneMatch.includes("github")){
            finalMatch.push({ "url": oneMatch, "account": "Github"})
          } else if (oneMatch.includes("dev.to")){
            finalMatch.push({ "url": oneMatch, "account": "Dev.to"})
          } else if (oneMatch.includes("bitbucket")){
            finalMatch.push({ "url": oneMatch, "account": "Bitbucket"})
          }
        });
        return finalMatch
      },

    },

    methods: {
      checkUsername() {
        this.loading = true
        axios.post(`${API_ROUTE}/username`, this.checkUrl)
             .then(res => {
               console.log("this is our response: ", res)
               if (res.data.length > 0) {
                 this.matchedAccount = res.data
                 this.notMatched = false
               } else {
                 this.notMatched = true
               }
               this.loading = false
             }).catch(err => {
              this.loading = false
              this.notMatched = false
              this.matchedAccount = []
              console.log("This is the error: ", err)
           })
        }
      }
  };
</script>


<style scoped>
  button {
    color: white!important;
  }
  a {
    text-decoration: none;
    cursor: pointer;
  }
  .style-name {
    color: white!important;
  }

</style>
