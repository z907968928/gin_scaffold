# /bin/bash
#init color

RED='\E[1;31m'
GREEN='\E[1;32m'
YELLOW='\E[1;33m'
BLUE='\E[1;34m'
RES='\E[0m'


echo "---------------------------------------------------------------------------------------"
echo "this gin scaffold project as: github.com/e421083458/gin_scaffold"
echo "when enter a new group & project, this script will replace [e421083458/gin_scaffold] => [group/project name] & use project name create new dir"
echo "---------------------------------------------------------------------------------------"
echo ""
echo ""

read -p "please enter you project group: " group
if [ ! -n $group ];then
  echo "you enter group empty"
  exit 0
fi

read -p "please enter you project name: " project
if [ ! -n $project ];then
  echo "you enter project empty"
  exit 0
fi


group=${group/\//\\/}
result=$(echo $project |grep '/')
if [[ $result != "" ]];then
  echo "project can not contain '/'"
  exit 0
fi


echo -e "you enter group as : ${RED} ${group} ${RES}"
echo -e "you enter project name as : ${RED} ${project} ${RES}"

echo "---------------------------------------------------------"
echo "create ${project} dir"
if [ -d project ];then
  echo "then ${project} already exist"
  exit 0
fi


mkdir ${project}
cp -r * $project
group=`echo ${group/\//\\/}`
echo "replace you project all [e421083458/gin_scaffold] => ${group}/${project}"
echo "cd ${project}  && sed -i \"s/e421083458\/gin_scaffold/${group}\/${project}/g\" \`grep -rl \"e421083458/gin_scaffold\`"
cd ${project} && sed -i "s/e421083458\/gin_scaffold/${group}\/${project}/g" `grep -rl "e421083458/gin_scaffold"`
