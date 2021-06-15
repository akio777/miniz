module.exports = (timestr=null) => {
    const date = timestr?(new Date(timestr)):(new Date())
    const dd = parseInt(date.getDate()) > 9 ? parseInt(date.getDate()) : '0' + parseInt(date.getDate())
    const mm = (parseInt(date.getMonth()) + 1) > 9 ? (parseInt(date.getMonth()) + 1) : '0' + (parseInt(date.getMonth()) + 1)
    const yyyy = ('' + date.getFullYear())
    const hh = parseInt(date.getHours()) > 9 ? parseInt(date.getHours()) : '0' + parseInt(date.getHours())
    const min = parseInt(date.getMinutes()) > 9 ? parseInt(date.getMinutes()) : '0' + parseInt(date.getMinutes())
    const sec = parseInt(date.getSeconds()) > 9 ? parseInt(date.getSeconds()) : '0' + parseInt(date.getSeconds())
    return yyyy + '' + mm + '' + dd + '_' + hh + '' + min +  '' + sec
  }