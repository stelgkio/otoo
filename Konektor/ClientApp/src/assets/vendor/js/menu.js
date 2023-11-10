function testMethodAdd() {

  if( document.getElementById('htmlId').offsetWidth<1199){
    document.getElementById("htmlId").classList.add('layout-menu-expanded');
  }else{
    document.getElementById("htmlId").classList.add('layout-menu-collapsed');
  }




}
function testMethodRemove() {
  if( document.getElementById('htmlId').offsetWidth <1199){
    document.getElementById("htmlId").classList.remove('layout-menu-expanded');
  }else{
    document.getElementById("htmlId").classList.remove('layout-menu-collapsed');
  }

}
