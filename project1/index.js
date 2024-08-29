const rule  = {
    AND: {
        Type:true,
        OR: {
            Description:true,
            RefLink:true
        }
    }
        // OR: {
        //     Description:true,
        //     RefLink:true
        // }
}
const obj={
    Type: {
        label: "Webinar",
        value: "WEBINAR"
    },
    Description: null,
    RefLink: null,
    key: "00fb491e-8ee2-43a5-859f-13dc7f6262ba",
    index: 0,
    id: "a5d3a9a5-0462-40cd-9843-20b37a8b93ed"
}
const parser = (rules, data) => {
  const [operator, value] = Object.entries(rules)[0];
  console.log("🚀 ===== parser ===== value:", value);
  console.log("🚀 ===== parser ===== operator:", operator);
  let result = null;
  if (operator === 'AND') {
    for (const item of Object.entries(value)) {
        const [key,_value] =item
        if(key!=='OR' && !Boolean(data?.[key])) return false
        if(key==='OR'){
            console.log('OR condition')
            const orResult = parser({[key]:_value}, data)
            result=result&&orResult
            continue
        }
        if(result===null){
            result=Boolean(data?.[key])
        }else{
            result=result && Boolean(data?.[key])
        }
    }
  }
  if(operator==='OR'){
    for(let item of Object.entries(value) ){
        const [key, _value] = item
        if(key==='AND'){
            result=result || parser({[key]:_value},data)
            continue
        }
        if(result===null){
            result=Boolean(data?.[key])
        }else{
            result=result || Boolean(data?.[key])
        }
    }
  }
  return result
};
console.log('parser',parser(rule,obj))
// console.log(Object.entries(rule))