create or replace  FUNCTION get_or_create_delivery(name varchar,phone varchar, zip varchar,city varchar,address varchar,region varchar, email varchar) returns integer 
as $$
declare id integer default  0;
begin
	select id_delivery into id 
	from delivery  
	where delivery.name = $1 and 
	delivery.phone= $2 and 
	delivery.zip=$3 and 
	delivery.city = $4 and 
	delivery.address = $5 and
	delivery.region = $6 and
	delivery.email =$7 ;
	if  found  then 
		return id;
	end if;
	insert into delivery (name,phone , zip ,city,address ,region , email) Values(name,phone , zip ,city,address ,region , email) returning id_delivery into id;
	return  id;
end;
$$ language plpgsql;
select 	get_or_create_delivery('sname','sphone', 'szip','scity' ,'saddress','sregion' , 'semail') 

drop function get_or_create_payment;
create or replace  FUNCTION get_or_create_payment(transaction varchar,request_id varchar, 
currensy varchar,provider varchar,amount integer ,paymont_dt integer,
bank varchar,delivery_cost integer, goods_total integer, custom_fee integer)
returns integer 
as $$
declare id integer default  0;
begin
	select id_payment into id 
	from payment  
	where payment.transaction  = $1 and 
	payment.request_id= $2 and 
	payment.currensy= $3 and 
	payment.provider = $4 and 
	payment.amount = $5 and
	payment.payment_dt = $6 and
	payment.bank = $7 and
	payment.delivery_cost = $8 and
	payment.goods_total = $9 and
	payment.custom_fee = $10;

	if  found  then 
		return id;
	end if;
	insert into payment (transaction,request_id, currensy ,provider,amount ,payment_dt, bank,delivery_cost,goods_total,custom_fee) 
		Values(transaction,request_id, currensy ,provider,amount ,paymont_dt, bank,delivery_cost,goods_total,custom_fee) 
		returning id_payment  into id;
	return  id;
end;
$$ language plpgsql;

Select get_or_create_payment('b563feb7b2b84b6test','','USD','wbpay',1817,1637907727,'alpha',1500,317,0);

drop function get_or_create_item;
create or replace  FUNCTION get_or_create_item(chrt_idv integer ,track_number varchar, price integer, 
rid varchar, name varchar, sale integer, size varchar,total_price integer, nm_id integer, brand varchar, status integer) 
returns integer 
as $$
declare id integer default  0;
begin
	select chrt_id  into id 
	from item 
	where item.chrt_id  = $1;
	if  found  then 
		return id;
	end if;
	insert into item(chrt_id ,track_number, price ,rid,name ,sale , size,total_price,nm_id,brand,status) Values(chrt_idv ,track_number , price ,rid,name ,sale , size,total_price,nm_id,brand,status) returning chrt_id  into id;
	return  id;
end;
$$ language plpgsql;



create or replace  procedure create_order(order_uid varchar,track_number varchar,entry varchar, 
delivery_id integer,payment_id integer, locale varchar,internal_signature varchar,
customer_id varchar,delivery_service varchar,shardkey varchar, sm_id integer,
date_created varchar,oof_shard varchar 
)  
language  sql 
as $$
	insert into orders(order_uid , track_number  ,entry,delivery_id ,payment_id , locale,internal_signature,customer_id,delivery_service,shardkey,sm_id,date_created,oof_shard )
Values( order_uid , track_number  ,entry,delivery_id ,payment_id , locale,internal_signature,customer_id,delivery_service,shardkey,sm_id,date_created,oof_shard );
$$ ;

Select get_or_create_item(99349310,'WBILMTESTTRACK',453,'ab4219087a764ae0btest','Mascaras',30,'0',317,2389212,'Vivienne Sabo',202);

create or replace  procedure  add_itemInDelivery(item_id integer,order_id varchar)
language  sql 
as $$
	insert into item_order (item_id,order_id) values(item_id,order_id)
$$;

create  view View_orders_ALL as 
	select * from orders;


CREATE OR REPLACE FUNCTION get_delivery_byId(id integer)
RETURNS SETOF delivery  AS $$
BEGIN
  RETURN QUERY 
  select *
 FROM delivery 
where id_delivery = id  ;
END
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION get_payment_byId(id integer)
RETURNS SETOF payment  AS $$
BEGIN
  RETURN QUERY 
  select *
 FROM payment 
where id_payment  = id  ;
END
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION get_items_byorderId(id varchar)
RETURNS SETOF item  AS $$
BEGIN
  RETURN QUERY 
  select * 
  from  item 
  where chrt_id in
  (select item_id 
  from item_order 
  where order_id = id
  );
	  
END
$$ LANGUAGE plpgsql;
select* from payment ;
Select get_or_create_delivery('Test Testov'::varchar,'+9720000000'::varchar,'2639809'::varchar,'Kiryat Mozkin'::varchar,'Ploshad Mira 15'::varchar,'Kraiot'::varchar,'test@gmail.com'::varchar)

Select get_delivery_byId(122);

select get_items_byorderId(' e21e21s1a')
	