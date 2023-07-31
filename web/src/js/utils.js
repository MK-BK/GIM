export default {
    formatTime(time) {
        let date = new Date(time)
        let result;
        const options = {
            "Y+": date.getFullYear().toString(),
            "m+": (date.getMonth() + 1).toString(),
            "d+": date.getDate().toString(),
            "H+": date.getHours().toString(),
            "M+": date.getMinutes().toString(),
            "S+": date.getSeconds().toString()
        };

        let format = "YYYY-mm-dd HH:MM"
        for (let option in options) {
            result = new RegExp("(" + option + ")").exec(format);
            if (result) {
                format = format.replace(result[1], (result[1].length == 1) ? (options[option]) : (options[option].padStart(result[1].length, '0')))
            };
        };
        return format;
    },
}