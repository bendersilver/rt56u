<template>
  <div id="app">
    <nav class="navbar navbar-expand-lg fixed-top navbar-light bg-light">
      <div class="container-fluid justify-content-center">
        <ul class="navbar-nav d-flex flex-row">
          <li class="nav-item me-3 me-lg-0 d-flex align-items-center ">
            <div class="form-check form-switch">
              <input class="form-check-input" type="checkbox" v-model="editHide" />
            </div>
          </li>
          <li class="nav-item me-3 me-lg-0">
            <a class="nav-link" href="#" @click="showHide = !showHide">
              <i :class="(showHide ? 'fa-eye-slash' : 'fa-eye') + ' fas'"></i>
            </a>
          </li>
          <li class="nav-item me-3 me-lg-0">
            <a
              class="nav-link"
              href="#"
              :class="{ active: sorting == 'namervs' }"
              @click="sortingSet('namervs')"
            >
              <i class="fas fa-sort-alpha-up"></i>
            </a>
          </li>
          <li class="nav-item me-3 me-lg-0">
            <a
              class="nav-link"
              href="#"
              :class="{ active: sorting == 'name' }"
              @click="sortingSet('name')"
            >
              <i class="fas fa-sort-alpha-down"></i>
            </a>
          </li>
          <li class="nav-item me-3 me-lg-0">
            <a
              class="nav-link"
              href="#"
              :class="{ active: sorting == 'orderrvs' }"
              @click="sortingSet('orderrvs')"
            >
              <i class="fas fa-sort-numeric-up"></i>
            </a>
          </li>
          <li class="nav-item me-3 me-lg-0">
            <a
              class="nav-link"
              href="#"
              :class="{ active: sorting == 'order' }"
              @click="sortingSet('order')"
            >
              <i class="fas fa-sort-numeric-down"></i>
            </a>
          </li>
          <li class="nav-item me-3 me-lg-0">
            <a
              class="nav-link"
              href="#"
              :class="{ active: sorting == 'group' }"
              @click="sortingSet('group')"
            >
              <i class="fas fa-sort-amount-down"></i>
            </a>
          </li>
          <li
            class="nav-item me-3 d-flex me-lg-0 align-items-center text-muted"
          >
            <small class="info">
              Всего: {{ list.length }} <i class="fas fa-eye"></i>
              {{ list.filter((i) => !i.del && !i.hide).length }}
              <i class="fas fa-eye-slash"></i>
              {{ list.filter((i) => !i.del && i.hide).length }}
              <i class="fas fa-trash"></i>
              {{ list.filter((i) => i.del).length }}
            </small>
          </li>
        </ul>
      </div>
    </nav>
    <div class="vh-100 overflow-hidden">
      <div class="overflow-auto body">
        <div class="container">
          <div class="row">
            <div class="col">
              <div class="w100 d-flex justify-content-center">
                <draggable
                  class="list-group list-group-flush overflow-auto"
                  tag="ul"
                  v-model="list"
                  v-bind="dragOptions"
                  handle=".fa-sort"
                  @change="udateList"
                >
                  <transition-group type="transition">
                    <li
                      class="list-group-item"
                      v-for="i in list"
                      :key="i.id"
                      v-show="showHide ? true : !i.hide"
                    >
                      <div class="row align-items-center">
                        <div class="col-3">
                          <img :src="i.img" alt="..." class="img-fluid" />
                        </div>
                        <div class="col-5">
                          {{ i.name }}<br />
                          <small class="text-muted">{{ i.group }}</small>
                        </div>
                        <div
                          class="col-4 d-flex justify-content-between action"
                        >
                          <i
                            v-if="editHide"
                            :class="
                              (i.hide ? 'fa-eye-slash' : 'fa-eye') + ' fas'
                            "
                            @click="hideToggle(i)"
                          ></i>
                            <input
                            v-model="i.tmpsprt"
                            :placeholder="i.order"
                              v-if="sorting == 'order'"
                              type="text"
                              class="form-control form-control-sm"
                            />
                            <i v-if="sorting == 'order'" class="fas fa-check" @click="serOrder(i, i.tmpsprt);i.tmpsprt=null"></i>
                          <i v-if="sorting == 'order'" class="fas fa-sort"></i>
                        </div>
                      </div>
                    </li>
                  </transition-group>
                </draggable>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div v-if="sorting == 'order'"  class="fixed-action-btn">
      <a class="btn btn-floating btn-primary" @click="saveList">
        <i class="fas fa-save"></i>
      </a>
    </div>
  </div>
</template>

<script>
import draggable from "vuedraggable";
export default {
  name: "App",
  components: {
    draggable,
  },
  mounted: function () {
    fetch("/jsonAPI/get")
      .then((response) => response.json())
      .then((data) => {
        this.list = data;
        this.sortingSet(this.sorting);
      });
  },
  data() {
    return {
      editHide: false,
      sorting: "order",
      showHide: false,
      list: [],
    };
  },
  methods: {
    sortingSet(v) {
      this.sorting = v;
      switch (v) {
        case "namervs":
          this.list = this.list.sort(
            (a, b) =>
              a.hide - b.hide ||
              b.name.toLowerCase().localeCompare(a.name.toLowerCase())
          );
          break;
        case "name":
          this.list = this.list.sort(
            (a, b) =>
              a.hide - b.hide ||
              a.name.toLowerCase().localeCompare(b.name.toLowerCase())
          );
          break;
        case "orderrvs":
          this.list = this.list.sort(
            (a, b) => a.hide - b.hide || b.order - a.order
          );
          break;
        case "order":
          this.list = this.list
            .sort((a, b) => a.hide - b.hide || a.order - b.order)
            .map((i, ix) => {
              i.order = ix;
              return i;
            });
          break;
        case "group":
          this.list = this.list.sort(
            (a, b) =>
              a.hide - b.hide ||
              a.group.toLowerCase().localeCompare(b.group.toLowerCase()) ||
              a.name.toLowerCase().localeCompare(b.name.toLowerCase())
          );
          break;
      }
    },
    hideToggle(i) {
      i.hide = !i.hide;
      fetch("/jsonAPI/toggle", {
        method: "POST",
        body: JSON.stringify({id:i.id, hide: i.hide}),
      }).then(rsp=> {
        if (rsp.status !=  200) i.hide = !i.hide
        else this.sortingSet(this.sorting)
      })
    },
    serOrder(i, ix) {
      let num = parseInt(ix)
      if (num >=0) {
        i.order = num
        this.sortingSet(this.sorting)
      }
    },
    saveList() {
      fetch("/jsonAPI/save", {
        method: "POST",
        body: JSON.stringify(this.list),
      });
    },
    udateList() {
      this.list = this.list
            .map((i, ix) => {
              i.order = ix;
              return i;
            });
    },
  },
  computed: {
    dragOptions() {
      return {
        animation: 200,
        group: "description",
        disabled: false,
        ghostClass: "ghost",
      };
    },
  },
};
</script>

<style>
/* .info > i {
  margin-left: 10px;
} */
.action > i {
  margin: 5px;
  cursor: pointer;
}

.fas.fa-sort {
  cursor: move;
}
.btn > .fas {
  font-size: 1.5em;
}
.nav.nav-tabs.position-fixed {
  z-index: 100;
  width: 100%;
}
.vh-100 {
  padding-top: 60px;
}
.body {
  height: 100%;
}

.list-group {
  max-width: 500px;
}
</style>
