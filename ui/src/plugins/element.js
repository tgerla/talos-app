import '../assets/theme/index.css'
import ElementUI from 'element-ui'
import Vue from 'vue'
import {
    Button,
    Tabs,
    TabPane,
    Table,
    TableColumn,
    Loading,
    Timeline,
    TimelineItem,
    Input,
    Drawer,
    Container,
    Aside,
    Header,
    Main,
    Radio,
    RadioButton,
    RadioGroup
} from 'element-ui'
import lang from 'element-ui/lib/locale/lang/en'
import locale from 'element-ui/lib/locale'

Vue.use(ElementUI)
locale.use(lang)

Vue.use(Button)
Vue.use(Tabs)
Vue.use(TabPane)
Vue.use(Table)
Vue.use(TableColumn)
Vue.use(Loading)
Vue.use(Timeline)
Vue.use(TimelineItem)
Vue.use(Input)
Vue.use(Drawer)
Vue.use(Container)
Vue.use(Aside)
Vue.use(Header)
Vue.use(Main)
Vue.use(Radio)
Vue.use(RadioButton)
Vue.use(RadioGroup)
