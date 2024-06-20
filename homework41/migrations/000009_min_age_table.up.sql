select id, name, age, email, is_working from users 
group by id, name, age, email, is_working
having age = min(age)